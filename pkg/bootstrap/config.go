package bootstrap

import (
	"app/pkg/config"

	"github.com/joho/godotenv"
)

type Config struct {
	Env      config.EnvConfig
	Server   config.ServerConfig
	Database config.DatabaseConfig
	Password config.PasswordConfig
}

/* perform application bootstrap */
func InitConfig() (Config, error) {
	if err := LoadEnvironment(config.Env); err != nil {
		return Config{}, err
	}

	serverConfig, err := config.ServerConfigInit()
	if err != nil {
		return Config{}, err
	}

	dbConfig, err := config.DatabaseConfigInit()
	if err != nil {
		return Config{}, err
	}

	config := Config{
		Env:      config.Env,
		Server:   serverConfig,
		Database: dbConfig,
		Password: config.Password,
	}

	return config, nil
}

/* read env variables from file and check if all mandatory ones are provided */
func LoadEnvironment(c config.EnvConfig) error {
	if err := godotenv.Load(c.File); err != nil {
		return err
	}

	if err := c.Validate(); err != nil {
		return err
	}

	return nil
}
