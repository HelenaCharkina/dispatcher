package service

import (
	"dispatcher/pkg/repository"
	"dispatcher/types"
)

type Authorization interface {
	SignIn(login string, password string) (*types.Response, error)
	RefreshToken(refreshToken string, userId string) (*types.Response, error)
	CheckToken(accessToken string) (string, error)
	Logout(userId string) error
}

type Agent interface {
	GetAll() (*[]types.Agent, error)
	Add(agent *types.Agent) error
	Update(agent *types.Agent) error
	Delete(agentId string) error
	SetState(agentId string, state types.State) error
}

type Statistics interface {
	Add(stats *types.Statistics) error
	GetByAgentId(params *types.StatisticsRequest) (*[]types.DBStatistics, error)
}

type Service struct {
	Authorization
	Agent
	Statistics
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.User, repo.Token),
		Agent:         NewAgentService(repo.Agent),
		Statistics:    NewStatisticsService(repo.Statistics),
	}
}
