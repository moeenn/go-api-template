package env

import (
	"errors"
	"os"
	"strings"
)

func Env(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("environment variable '" + key + "' not set")
	}

	return strings.TrimSpace(value), nil
}
