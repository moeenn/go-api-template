package auth

import (
	"time"
	"web/internal/helpers/jwt"
)

const (
	SCOPE_ACCESS  = "ACCESS"
	SCOPE_REFRESH = "REFRESH"
)

type LoginAuthTokenPayload struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expiry       int64  `json:"expiry"`
}

// NewLoginAuthTokenPayload issue the jwt access and refresh token combo together
// with the expiration timestamp for the access token.
func NewLoginAuthTokenPayload(secret string, expiry time.Duration, user jwt.JWTUser) (LoginAuthTokenPayload, error) {
	accessToken, err := jwt.NewExpiringJWT(secret, expiry, user, SCOPE_ACCESS)
	if err != nil {
		return LoginAuthTokenPayload{}, err
	}

	refreshToken, err := jwt.NewNonExpiringJWT(secret, user, SCOPE_REFRESH)
	if err != nil {
		return LoginAuthTokenPayload{}, err
	}

	return LoginAuthTokenPayload{
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
		Expiry:       accessToken.Expiry,
	}, nil
}
