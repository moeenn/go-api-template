package jwt

import (
	"testing"
	"time"
)

var (
	user = JWTUser{
		Id:   "A100",
		Role: "ADMIN",
	}
	secret = "abc12312313"
	scope  = "ACCESS"
)

func TestCreateAndValidateToken(t *testing.T) {
	token, err := NewExpiringJWT(secret, time.Minute, user, scope)
	if err != nil {
		t.Errorf("token creation failed: %s", err.Error())
		return
	}

	gotUser, err := ValidateJWT(secret, token.Token, true, scope)
	if err != nil {
		t.Errorf("failed to validate valid token: %s", err.Error())
		return
	}

	if gotUser.Id != user.Id || gotUser.Role != user.Role {
		t.Errorf("unexpected user data receievd when validating token user: %+v", gotUser)
		return
	}
}

func TestCreateAndValidateNonExpiringToken(t *testing.T) {
	token, err := NewNonExpiringJWT(secret, user, scope)
	if err != nil {
		t.Errorf("token creation failed: %s", err.Error())
		return
	}

	gotUser, err := ValidateJWT(secret, token.Token, false, scope)
	if err != nil {
		t.Errorf("failed to validate valid token: %s", err.Error())
		return
	}

	if gotUser.Id != user.Id || gotUser.Role != user.Role {
		t.Errorf("unexpected user data received when validating token user: %+v", gotUser)
		return
	}
}

func TestCreateAndValidateExpiredToken(t *testing.T) {
	token, err := NewExpiringJWT(secret, time.Millisecond*20, user, scope)
	if err != nil {
		t.Errorf("token creation failed: %s", err.Error())
		return
	}

	time.Sleep(time.Millisecond * 50)
	user, err = ValidateJWT(secret, token.Token, true, scope)

	if err == nil {
		t.Errorf("expired token validated successfully: User %+v", user)
		return
	}
}

func TestCreateAndValidateInvalidToken(t *testing.T) {
	token, err := NewExpiringJWT(secret, time.Millisecond*20, user, scope)
	if err != nil {
		t.Errorf("token creation failed: %s", err.Error())
		return
	}

	user, err = ValidateJWT(secret, token.Token+"12312313", true, scope)
	if err == nil {
		t.Errorf("invalid token validated successfully: User %+v", user)
		return
	}
}

func TestCreateAndValidateInvalidScopeToken(t *testing.T) {
	token, err := NewExpiringJWT(secret, time.Millisecond*20, user, scope)
	if err != nil {
		t.Errorf("token creation failed: %s", err.Error())
		return
	}

	user, err = ValidateJWT(secret, token.Token+"12312313", true, "REFRESH")
	if err == nil {
		t.Errorf("invalid scoped token validated successfully: User %+v", user)
		return
	}
}
