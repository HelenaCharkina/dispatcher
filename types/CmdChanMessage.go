package types

type Cmd byte

const (
	START = 1
	STOP  = 2
)

type CmdChanMessage struct {
	Message Cmd
	AgentId string
}
