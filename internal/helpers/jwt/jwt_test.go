package jwt

import (
	"testing"
	"time"
)

var (
	secret = "abc123123123"
	user   = JWTUser{
		Id:   "A100",
		Role: "ADMIN",
	}
)

func TestJWTCreateValidate(t *testing.T) {
	token, err := NewJWT(secret, time.Minute, user)
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

func TestJWTExpiredValidation(t *testing.T) {
	token, _ := NewJWT(secret, time.Millisecond*20, user)
	time.Sleep(time.Millisecond * 30)

	_, err := ValidateJWT(secret, token.Token)
	if err == nil {
		t.Errorf("error not returned during verification of expired JWT")
	}
}
