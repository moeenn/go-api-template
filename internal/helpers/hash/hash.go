package hash

import (
	"github.com/alexedwards/argon2id"
)

func NewHash(clearText string) (string, error) {
	return argon2id.CreateHash(clearText, argon2id.DefaultParams)
}

func VerifyHash(clearText, hashed string) bool {
	match, err := argon2id.ComparePasswordAndHash(clearText, hashed)
	if err != nil {
		return false
	}

	return match
}
