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
		stmt, _ = tx.Prepare("INSERT INTO monitoring.agents (id, ip, port, name, schedule, description, state) VALUES (?, ?, ?, ?, ?, ?, 1)")
	)
	defer stmt.Close()

	id, err := uuid.NewV4()
	if _, err = stmt.Exec(
		id,
		agent.Ip,
		agent.Port,
		agent.Name,
		agent.Schedule,
		agent.Description,
	); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
func (r *AgentRepo) Update(agent *types.Agent) error {
	return nil
}
func (r *AgentRepo) Delete(agentId string) error {
	return nil
}
