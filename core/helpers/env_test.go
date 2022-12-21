package helpers

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	os.Setenv("example", "sample")

	value, err := Env("example")
	if err != nil {
		t.Errorf("unexpected error thrown")
	}

	if value != "sample" {
		t.Errorf("unexpected env value received: %s", value)
	}
}
