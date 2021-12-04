package repository

import (
	"dispatcher/types"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetUser(login, password string) (*types.User, error)
}

type Agent interface {
	GetAll() (*[]types.Agent, error)
	Add(agent *types.Agent) error
	Update(agent *types.Agent) error
	Delete(agentId string) error
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
}

func NewRepository(db *sqlx.DB, cache *redis.Client) *Repository {
	return &Repository{
		User:  NewUserRepo(db),
		Token: NewTokenRepo(cache),
		Agent: NewAgentRepo(db),
	}
}
