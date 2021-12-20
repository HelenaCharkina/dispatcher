package repository

import (
	"dispatcher/types"
	"github.com/jmoiron/sqlx"
	uuid "github.com/nu7hatch/gouuid"
)

type AgentRepo struct {
	db *sqlx.DB
}

func NewAgentRepo(db *sqlx.DB) *AgentRepo {
	return &AgentRepo{
		db: db,
	}
}

func (r *AgentRepo) GetAll() (*[]types.Agent, error) {
	var agents []types.Agent
	rows, err := r.db.Query("SELECT id, ip, port, name, schedule, description, state FROM monitoring.agents")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var agent types.Agent
		if err = rows.Scan(&agent.Id, &agent.Ip, &agent.Port, &agent.Name, &agent.Schedule, &agent.Description, &agent.State); err != nil {
			return nil, err
		}
		agents = append(agents, agent)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &agents, nil
}
func (r *AgentRepo) Add(agent *types.Agent) error {
	var (
		tx, _   = r.db.Begin()
		stmt, _ = tx.Prepare("INSERT INTO monitoring.agents (id, ip, port, name, schedule, description, state) VALUES (?, ?, ?, ?, ?, ?, ?)")
	)
	defer stmt.Close()

	id, err := uuid.NewV4()
	if _, err = stmt.Exec(
		id.String(),
		agent.Ip,
		agent.Port,
		agent.Name,
		agent.Schedule,
		agent.Description,
		1,
	); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
func (r *AgentRepo) Update(agent *types.Agent) error {

	var (
		tx, _   = r.db.Begin()
		stmt, _ = tx.Prepare("alter table monitoring.agents update ip=?, port=?, name=?, description=?, schedule=? where id=?")
	)
	defer stmt.Close()

	if _, err := stmt.Exec(
		agent.Ip,
		agent.Port,
		agent.Name,
		agent.Description,
		agent.Schedule,
		agent.Id,
	); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
func (r *AgentRepo) Delete(agentId string) error {
	var (
		tx, _   = r.db.Begin()
		stmt, _ = tx.Prepare("alter table monitoring.agents delete where id=?")
	)
	defer stmt.Close()


	if _, err := stmt.Exec(
		agentId,
	); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *AgentRepo) SetState(agentId string, state types.State) error  {
	var (
		tx, _   = r.db.Begin()
		stmt, _ = tx.Prepare("alter table monitoring.agents update state=? where id=?")
	)
	defer stmt.Close()

	if _, err := stmt.Exec(
		state,
		agentId,
	); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}