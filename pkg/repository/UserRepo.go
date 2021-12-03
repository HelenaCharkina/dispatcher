package repository

import (
	"dispatcher/types"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetUser(login, password string) (*types.User, error) {
	var user types.User
	query := "select id, name from users where login = ? and password = ?"
	err := r.db.Get(&user, query, login, password)

	return &user, err
}
