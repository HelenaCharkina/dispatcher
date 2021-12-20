package repository

import (
	"dispatcher/types"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetUser(login, password string) (*types.User, error)
}

type Statistics interface {
	Add(stats *types.Statistics) error
	GetByAgentId(params *types.StatisticsRequest) (*[]types.DBStatistics, error)
}


type Agent interface {
	GetAll() (*[]types.Agent, error)
	Add(agent *types.Agent) error
	Update(agent *types.Agent) error
	Delete(agentId string) error
	SetState(agentId string, state types.State) error
}

type Token interface {
	SetToken(token string, userId string) error
	GetToken(userId string) (string, error)
	RemoveToken(userId string) error
}

type Repository struct {
	User
	Agent
	Token
	Statistics
}

func NewRepository(db *sqlx.DB, cache *redis.Client) *Repository {
	return &Repository{
		User:       NewUserRepo(db),
		Token:      NewTokenRepo(cache),
		Agent:      NewAgentRepo(db),
		Statistics: NewStatisticsRepo(db),
	}
}
