package repository

import (
	"dispatcher/types"
	"github.com/jmoiron/sqlx"
)

type AuthClickhouse struct {
	db *sqlx.DB
}

func NewAuthClickhouse(db *sqlx.DB) *AuthClickhouse {
	return &AuthClickhouse{
		db: db,
	}
}

func (r *AuthClickhouse) GetUser(login, password string) (*types.User, error) {
	var user types.User
	query := "select id, name from users where login = ? and password = ?"
	err := r.db.Get(&user, query, login, password)

	return &user, err
}
