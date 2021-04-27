package config

import (
	"github.com/kelseyhightower/envconfig"
)

type environment struct {
	DBHost string `envconfig:"db_host"`                // SENZR_DB_HOST
	DBName string `envconfig:"db_name"`                // SENZR_DB_NAME
	DBPort string `envconfig:"db_port" default:"5432"` // SENZR_DB_PORT
	DBUser string `envconfig:"db_user"`                // SENZR_DB_USER
}

func Env() (*environment, error) {
	var e environment
	err := envconfig.Process("senzr", &e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
