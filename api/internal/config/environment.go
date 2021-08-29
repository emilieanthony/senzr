package config

import (
	"github.com/kelseyhightower/envconfig"
)

type environment struct {
	Project string `envconfig:"project" default:"senzr-313218"` // SENZR_PROJECT
	AuthKey string `envconfig:"auth_key"`                       // SENZR_AUTH_KEY
}

func Env() (*environment, error) {
	var e environment
	err := envconfig.Process("senzr", &e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
