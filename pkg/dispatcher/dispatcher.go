package dispatcher

import (
	"dispatcher/pkg/service"
	"dispatcher/pkg/settings"
	"dispatcher/types"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
	"time"
)

type mapElem struct {
	agent types.Agent
	ch    chan types.CmdChanMessage
}

type Dispatcher struct {
	service    *service.Service
	wsChan     chan<- types.WsChanMessage
	agentTable map[string]mapElem
	CmdChan    chan types.CmdChanMessage
}

func NewDispatcher(service *service.Service, wsChan chan<- types.WsChanMessage) *Dispatcher {
	return &Dispatcher{
		service:    service,
		wsChan:     wsChan,
		agentTable: make(map[string]mapElem),
		CmdChan:    make(chan types.CmdChanMessage),
	}
}

func (d *Dispatcher) Run() error {

	agents, err := d.service.Agent.GetAll()
	if err != nil {
		return err
	}

	for _, a := range *agents {
		a := a
		cmdChan := make(chan types.CmdChanMessage)
		d.agentTable[a.Id] = mapElem{
			a, cmdChan,
		}
		if a.State == types.STARTED {
			go d.runAgent(a, cmdChan)
		}
	}
	go func() {
		for {
			select {
			case msg := <-d.CmdChan:
				agent, ok := d.agentTable[msg.AgentId]
				if ok {
					if msg.Message == types.STOP {
							agent.ch <- msg
					} else if msg.Message == types.START {
						err = d.service.Agent.SetState(msg.AgentId, types.STARTED)
						if err != nil {
							d.sendError(agent.agent.Name, fmt.Sprintf("Ошибка при запуске агента: %s", err))
							continue
						}
						go d.runAgent(agent.agent, agent.ch)
					}
				}
			}
		}
	}()

	return nil
}

func (d *Dispatcher) runAgent(agent types.Agent, cmdChan chan types.CmdChanMessage) {
	timeout, err := d.getTimeout(agent.Schedule)
	if err != nil {
		logrus.Errorf("agent %s error: %s", agent.Id, err)
		return
	}

	nextRequest := time.After(0)
	for {
		select {
		case <-nextRequest:
			response, err := d.sendRequest(agent)
			if err != nil {
				err = d.service.Agent.SetState(agent.Id, types.ERROR)
				if err != nil {
					logrus.Errorf("Set state error: %s", err)
					continue
				}
			}
			if response != nil {
				d.save(response, agent.Id)
			}
			nextRequest = time.After(timeout)
		case msg := <-cmdChan:
			if agent.Id != msg.AgentId {
				continue
			}
			if msg.Message == types.STOP {
				err = d.service.Agent.SetState(msg.AgentId, types.STOPPED)
				if err != nil {
					d.sendError(agent.Name, fmt.Sprintf("Ошибка при остановке агента: %s", err))
					continue
				}
				return
			}
		}
	}
}

func (d *Dispatcher) getTimeout(schedule string) (time.Duration, error) {
	if len(schedule) < 2 {
		return 0, errors.New("Invalid schedule ")
	}
	t := schedule[:len(schedule)-1]
	ts, err := strconv.Atoi(t)
	if err != nil {
		return 0, fmt.Errorf("Convertation error: %s ", err)
	}

	switch schedule[len(schedule)-1:] {
	case "s":
		return time.Duration(ts) * time.Second, nil
	case "m":
		return time.Duration(ts) * time.Minute, nil
	case "h":
		return time.Duration(ts) * time.Hour, nil
	default:
		return 0, errors.New("Invalid schedule ")
	}
}

func (d *Dispatcher) sendRequest(agent types.Agent) ([]byte, error) {

	raddr, err := net.ResolveUDPAddr("udp4", net.JoinHostPort(agent.Ip, agent.Port))
	if err != nil {
		logrus.Errorf("ResolveUDPAddr error: %s", err)
		return nil, nil
	}

	conn, err := net.DialUDP("udp4", nil, raddr)
	if err != nil {
		logrus.Errorf("ResolveUDPAddr error: %s", err)
		return nil, nil
	}
	defer conn.Close()

	deadline := time.Now().Add(time.Duration(settings.Config.RequestTimeout) * time.Second)
	err = conn.SetDeadline(deadline)
	if err != nil {
		logrus.Errorf("SetReadDeadline error: %s", err)
		return nil, nil
	}

	request := []byte("get")
	_, err = conn.Write(request)
	if err != nil {
		logrus.Errorf("Write error: %s", err)
		d.sendError(agent.Name, "")
		return nil, err
	}

	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		logrus.Errorf("ReadFromUDP error: %s", err)
		d.sendError(agent.Name, "")
		return nil, err
	}

	return buffer[:n], nil
}

func (d *Dispatcher) sendError(agentName string, e string) {

	select {
	case d.wsChan <- types.WsChanMessage{
		Message: fmt.Sprintf("Агент %s недоступен. %s", agentName, e),
	}:
	default:
	}
}

func (d *Dispatcher) save(response []byte, agentId string) {
	var stats types.Statistics

	err := json.Unmarshal(response, &stats)
	if err != nil {
		logrus.Errorf("json.Unmarshal error: %s", err)
		return
	}

	stats.AgentId = agentId
	err = d.service.Statistics.Add(&stats)
	if err != nil {
		logrus.Errorf("Add stats error: %s", err)
		return
	}
}
