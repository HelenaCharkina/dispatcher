package repository

import (
	"dispatcher/types"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	GetUser(login, password string) (*types.User, error)
}

type Agent interface {

}

type Repository struct {
	Authorization
	Agent
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthClickhouse(db),
	}
}
