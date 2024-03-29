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

type DatabaseConfig struct {
	ConnectionString string
}

type AuthConfig struct {
	JWTSecret      string
	JWTExpiryHours time.Duration
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Auth     AuthConfig
}

func NewConfig() (*Config, error) {
	jwtSecret, err := env.Env("JWT_SECRET")
	if err != nil {
		return &Config{}, err
	}

	connString, err := env.Env("DB_CONNECTION")
	if err != nil {
		return &Config{}, err
	}

	config := &Config{
		Server: ServerConfig{
			Host: "0.0.0.0",
			Port: 5000,
		},
		Database: DatabaseConfig{
			ConnectionString: connString,
		},
		Auth: AuthConfig{
			JWTSecret:      jwtSecret,
			JWTExpiryHours: time.Hour,
		},
	}

	return config, nil
}
