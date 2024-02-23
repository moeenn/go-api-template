package request

import (
	"errors"
	"net/http"
	"web/internal/helpers/jwt"
)

var authError = errors.New("please login to continue")

const (
	USER_ID_CTX_KEY   = "user_id"
	USER_ROLE_CTX_KEY = "user_role"
)

func CurrentUser(r *http.Request) (jwt.JWTUser, error) {
	userId, ok := r.Context().Value(USER_ID_CTX_KEY).(string)
	if !ok {
		return jwt.JWTUser{}, authError
	}

	userRole, ok := r.Context().Value(USER_ROLE_CTX_KEY).(string)
	if !ok {
		return jwt.JWTUser{}, authError
	}

	user := jwt.JWTUser{
		Id:   userId,
		Role: userRole,
	}

	return user, nil
}
