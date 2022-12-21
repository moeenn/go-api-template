package config

import (
	"app/core/helpers"
	"fmt"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

/* TODO: @security - enable SSL mode */
func (d DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Database,
	)
}

func DatabaseConfigInit() (DatabaseConfig, error) {
	host, err := helpers.Env("DB_HOST")
	if err != nil {
		return DatabaseConfig{}, err
	}

	port, err := helpers.Env("DB_PORT")
	if err != nil {
		return DatabaseConfig{}, err
	}

	user, err := helpers.Env("DB_USER")
	if err != nil {
		return DatabaseConfig{}, err
	}

	pass, err := helpers.Env("DB_PASSWORD")
	if err != nil {
		return DatabaseConfig{}, err
	}

	db, err := helpers.Env("DB_DATABASE")
	if err != nil {
		return DatabaseConfig{}, err
	}

	config := DatabaseConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: pass,
		Database: db,
	}

	return config, nil
}
