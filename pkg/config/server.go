package config

import (
	"app/core/helpers"
)

type ServerConfig struct {
	Port           string
	TrustedProxies []string
}

func ServerConfigInit() (ServerConfig, error) {
	port, err := helpers.Env("SERVER_PORT")
	if err != nil {
		return ServerConfig{}, err
	}

	config := ServerConfig{
		Port: ":" + port,
		TrustedProxies: []string{
			"192.168.1.1",
		},
	}

	return config, nil
}
