package jwt

import (
	"testing"
)

func TestJWTCreateValidate(t *testing.T) {
	secret := "abc123123123"
	user := JWTUser{
		Id:   "A100",
		Role: "ADMIN",
	}

	token, err := NewJWT(secret, user)
	if err != nil {
		t.Errorf("token creation failed. Error: %s", err.Error())
	}

	validated, err := ValidateJWT(secret, token.Token)
	if err != nil {
		t.Errorf("validation failed to a valid token. Error: %s", err.Error())
	}

	if validated.Id != user.Id || validated.Role != user.Role {
		t.Errorf("unexpected data extracted from token: %v", validated)
	}
}
