package config

import (
	"fmt"
	"time"
	"web/internal/helpers/env"
)

type ServerConfig struct {
	Host string
	Port int
}

func (c *ServerConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type AuthConfig struct {
	JWTSecret      string
	JWTExpiryHours time.Duration
}

type Config struct {
	Server ServerConfig
	Auth   AuthConfig
}

func NewConfig() (*Config, error) {
	jwtSecret, err := env.Env("JWT_SECRET")
	if err != nil {
		return &Config{}, err
	}

	config := &Config{
		Server: ServerConfig{
			Host: "0.0.0.0",
			Port: 5000,
		},
		Auth: AuthConfig{
			JWTSecret:      jwtSecret,
			JWTExpiryHours: 24,
		},
	}

	return config, nil
}
