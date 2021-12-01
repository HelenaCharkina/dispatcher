package repository

import (
	"dispatcher/pkg/settings"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

func NewClickhouseDB() (*sqlx.DB, error) {
	conn, err := sqlx.Open("clickhouse",
		fmt.Sprintf("tcp://%s:%s?database=%s&debug=%s", settings.Config.DBHost, settings.Config.DBPort, settings.Config.DBName, settings.Config.DebugMode))
	if err != nil {
		return nil, fmt.Errorf("Connection open error: %v ", err)
	}
	if err := conn.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("clickhouse error [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, fmt.Errorf("Connection ping error: %v ", err)
	}
	return conn, nil
}
