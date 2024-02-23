package config

import (
	"fmt"
)

type ServerConfig struct {
	Host string
	Port int
}

func (c *ServerConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type AuthConfig struct {
	JWTSecret string
}

type Config struct {
	Server ServerConfig
	Auth   AuthConfig
}

func NewConfig() (*Config, error) {
	config := &Config{
		Server: ServerConfig{
			Host: "0.0.0.0",
			Port: 5000,
		},
		Auth: AuthConfig{
			JWTSecret: "abc123123123", // TODO: read from env
		},
	}

	return config, nil
}
