package config

import (
	"errors"
	"os"
)

type EnvConfig struct {
	File      string
	Mandatory []string
}

/**
 *  check if all the required environment variables have been set properly
 *
 */
func (e EnvConfig) Validate() error {
	for _, entry := range e.Mandatory {
		value := os.Getenv(entry)
		if value == "" {
			return errors.New("failed to read mandatory env variable: " + entry)
		}
	}

	return nil
}

func envConfigInit() EnvConfig {
	return EnvConfig{
		File: ".env.local",
		Mandatory: []string{
			"SERVER_PORT",
			"DB_HOST",
			"DB_PORT",
			"DB_USER",
			"DB_PASSWORD",
			"DB_DATABASE",
		},
	}
}

var Env EnvConfig = envConfigInit()
