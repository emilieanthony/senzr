package db

import (
	"fmt"

	"github.com/emilieanthony/senzr/internal/config"
	"github.com/jmoiron/sqlx"
)

const (
	DRIVER = "postgres"
)

func newDBConnection() (_ *sqlx.DB, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("connection error: %v", err)
		}
	}()
	env, err := config.Env()
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Connect(
		DRIVER,
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
			env.DBHost, env.DBPort, env.DBUser, env.DBName,
		),
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
