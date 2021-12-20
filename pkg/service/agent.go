package service

import (
	"dispatcher/pkg/repository"
	"dispatcher/types"
)

type AgentService struct {
	repo  repository.Agent
}

func NewAgentService(repo repository.Agent) *AgentService {
	return &AgentService{
		repo:  repo,
	}
}

func (s *AgentService) GetAll() (*[]types.Agent, error) {
	return s.repo.GetAll()
}
func (s *AgentService) Add(agent *types.Agent) error {
	return s.repo.Add(agent)
}
func (s *AgentService) Update(agent *types.Agent) error {
	return s.repo.Update(agent)
}
func (s *AgentService) Delete(agentId string) error {
	return s.repo.Delete(agentId)
}

func (s *AgentService) SetState(agentId string, state types.State) error {
	return s.repo.SetState(agentId, state)
}