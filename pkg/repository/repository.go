package repository

import (
	"dispatcher/types"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	GetUser(login, password string) (*types.User, error)
}

type Agent interface {
}

type Token interface {
	SetToken(token string, userId string) error
	GetToken(userId string) (string, error)
}

type Repository struct {
	Authorization
	Agent
	Token
}

func NewRepository(db *sqlx.DB, cache *redis.Client) *Repository {
	return &Repository{
		Authorization: NewAuthClickhouse(db),
		Token:         NewTokenCache(cache),
	}
}
