package helpers

import (
	"errors"
	"os"
)

func Env(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return value, errors.New("env variable not set: " + key)
	}

	return value, nil
}
