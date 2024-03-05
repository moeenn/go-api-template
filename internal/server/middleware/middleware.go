package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"web/internal/helpers/jwt"
	"web/internal/helpers/request"
	"web/internal/helpers/response"
)

const (
	BEARER_TOKEN_PREFIX = "Bearer "
)

// extract bearer token value from request authorization header
func parseBearerToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return header, errors.New("missing authorization bearer token")
	}

	if !strings.HasPrefix(header, BEARER_TOKEN_PREFIX) {
		return "", errors.New("only bearer tokens are supported for authorization")
	}

	token := strings.Replace(header, BEARER_TOKEN_PREFIX, "", 1)
	if len(token) == 0 {
		return "", errors.New("please provide a bearer token for authorization")
	}

	return token, nil
}

// ensure that only logged-in users can access an provided request handler
func LoggedInMiddleware(jwtSecret string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := parseBearerToken(r)
		if err != nil {
			response.SendErr(w, http.StatusUnauthorized, err.Error())
			return
		}

		user, err := jwt.ValidateJWT(jwtSecret, token, true, "ACCESS")
		if err != nil {
			response.SendErr(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), request.USER_ID_CTX_KEY, user.Id)
		ctx = context.WithValue(ctx, request.USER_ROLE_CTX_KEY, user.Role)
		next(w, r.WithContext(ctx))
	}
}
