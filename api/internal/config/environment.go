package config

import (
	"github.com/kelseyhightower/envconfig"
)

type environment struct {
	DBHost     string `envconfig:"db_host" default:"localhost"`  // SENZR_DB_HOST
	DBName     string `envconfig:"db_name" default:"senzr_db"`   // SENZR_DB_NAME
	DBPort     string `envconfig:"db_port" default:"5432"`       // SENZR_DB_PORT
	DBUser     string `envconfig:"db_user" default:"root"`       // SENZR_DB_USER
	DBPassword string `envconfig:"db_password" default:"gbgftw"` // SENZR_DB_PASSWORD
}

func Env() (*environment, error) {
	var e environment
	err := envconfig.Process("senzr", &e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
