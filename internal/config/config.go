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

type Config struct {
	Server ServerConfig
}

func NewConfig() (*Config, error) {
	config := &Config{
		Server: ServerConfig{
			Host: "0.0.0.0",
			Port: 5000,
		},
	}

	return config, nil
}
