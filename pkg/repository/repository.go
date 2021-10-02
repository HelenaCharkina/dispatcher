package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {

}

type Agent interface {

}

type Repository struct {
	Authorization
	Agent
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
