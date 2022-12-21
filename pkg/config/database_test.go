package config

import (
	"testing"
)

func TestConnectionString(t *testing.T) {
	config := DatabaseConfig{
		Host:     "192.168.1.1",
		Port:     "5555",
		User:     "user",
		Password: "pass",
		Database: "db",
	}

	expected := "postgresql://user:pass@192.168.1.1:5555/db?sslmode=disable"
	got := config.ConnectionString()

	if expected != got {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}
