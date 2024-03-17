package hash

import (
	"testing"
)

func TestValidHash(t *testing.T) {
	clearText := "some_random_password_123"
	hash, err := NewHash(clearText)
	if err != nil {
		t.Errorf(
			"password hashing failed: %s", err.Error(),
		)
	}

	isMatch := VerifyHash(clearText, hash)
	if !isMatch {
		t.Error("valid password verification failed")
	}
}

func TestInvalidHash(t *testing.T) {
	clearText := "some_random_password_123"
	hash, _ := NewHash(clearText)

	isMatch := VerifyHash(clearText+"111", hash)
	if isMatch {
		t.Error("invalid password matched during valification")
	}
}
