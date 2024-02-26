package jwt

import (
	"errors"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

var (
	invalidExpiredErr = errors.New("invalid or expired JWT provided")
	invalidClaimsErr  = errors.New("failed to parse JWT claims")
)

type JWTWithExpiry struct {
	Token  string `json:"token"`
	Expiry int64  `json:"expiry"`
}

type JWTUser struct {
	Id   string `json:"id"`
	Role string `json:"role"`
}

/* create JWT claim and sign using JWT secret */
func NewJWT(secret string, expiryHours time.Duration, user JWTUser) (JWTWithExpiry, error) {
	exp := time.Now().Add(time.Hour * expiryHours)
	claims := &jwtlib.RegisteredClaims{
		Subject:   user.Id,
		Issuer:    user.Role,
		ExpiresAt: jwtlib.NewNumericDate(exp),
	}

	t, err := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return JWTWithExpiry{}, err
	}

	return JWTWithExpiry{
		Token:  t,
		Expiry: exp.Unix(),
	}, nil
}

/* validate token and extract custom claims */
func ValidateJWT(secret string, tokenString string) (JWTUser, error) {
	parsed, err := jwtlib.Parse(tokenString, func(token *jwtlib.Token) (any, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return JWTUser{}, err
	}

	if !parsed.Valid {
		return JWTUser{}, invalidExpiredErr
	}

	id, idErr := parsed.Claims.GetSubject()
	role, roleErr := parsed.Claims.GetIssuer()
	expiry, expErr := parsed.Claims.GetExpirationTime()

	if idErr != nil || roleErr != nil || expErr != nil {
		return JWTUser{}, invalidClaimsErr
	}

	if expiry.Before(time.Now()) {
		return JWTUser{}, invalidExpiredErr
	}

	return JWTUser{
		Id:   id,
		Role: role,
	}, nil
}
