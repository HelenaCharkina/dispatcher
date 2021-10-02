package repository

import (
	"dispatcher/pkg/settings"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

func NewPostgresDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		settings.Config.DBHost,
		settings.Config.DBPort,
		settings.Config.DBUsername,
		settings.Config.DBName,
		os.Getenv("DB_PASSWORD"),
		settings.Config.DBSSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
