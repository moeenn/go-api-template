package request

import (
	"errors"
	"net/http"
	"web/internal/helpers/jwt"
)

var errAuthRequired = errors.New("please login to continue")

type (
	UserIdCtxKey   struct{}
	UserRoleCtxKey struct{}
)

func CurrentUser(r *http.Request) (jwt.JWTUser, error) {
	userId, ok := r.Context().Value(UserIdCtxKey{}).(string)
	if !ok {
		return jwt.JWTUser{}, errAuthRequired
	}

	userRole, ok := r.Context().Value(UserRoleCtxKey{}).(string)
	if !ok {
		return jwt.JWTUser{}, errAuthRequired
	}

	user := jwt.JWTUser{
		Id:   userId,
		Role: userRole,
	}

	return user, nil
}
