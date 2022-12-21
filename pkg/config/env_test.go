package config

import (
	"os"
	"testing"
)

func populateMandatoryEnvVariables(mandatory []string) {
	for _, v := range mandatory {
		os.Setenv(v, "random_value")
	}
}

func TestValidateShouldThrowError(t *testing.T) {
	if err := Env.Validate(); err == nil {
		t.Errorf("no error thrown when error was expected")
	}
}

func TestValidate(t *testing.T) {
	populateMandatoryEnvVariables(Env.Mandatory)

	if err := Env.Validate(); err != nil {
		t.Errorf("not all mandatory env variables are set")
	}
}
