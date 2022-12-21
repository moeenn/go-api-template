package config

import (
	"app/core/helpers/password"
)

type PasswordConfig struct {
	Argon2 password.Params
}

func passwordConfigInit() PasswordConfig {
	return PasswordConfig{
		Argon2: password.Params{
			Memory:      64 * 1024,
			Iterations:  3,
			Parallelism: 2,
			SaltLength:  16,
			KeyLength:   32,
		},
	}
}

var Password = passwordConfigInit()
