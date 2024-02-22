package jwt

import (
	"errors"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

const JWT_EXPIRY_HOURS = 24

var (
	INVALID_EXPIRED_ERR = errors.New("invalid or expired JWT provided")
	INVALID_CLAIMS      = errors.New("failed to parse JWT claims")
)

type JWTWithExpiry struct {
	Token  string `json:"token"`
	Expiry int64  `json:"expiry"`
}

type JWTUser struct {
	Id   string
	Role string
}

/* create JWT claim and sign using JWT secret */
func NewJWT(secret string, user JWTUser) (JWTWithExpiry, error) {
	exp := time.Now().Add(time.Hour * JWT_EXPIRY_HOURS)
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
		return JWTUser{}, INVALID_EXPIRED_ERR
	}

	id, idErr := parsed.Claims.GetSubject()
	role, roleErr := parsed.Claims.GetIssuer()
	expiry, expErr := parsed.Claims.GetExpirationTime()

	if idErr != nil || roleErr != nil || expErr != nil {
		return JWTUser{}, INVALID_CLAIMS
	}

	if expiry.Before(time.Now()) {
		return JWTUser{}, INVALID_EXPIRED_ERR
	}

	return JWTUser{
		Id:   id,
		Role: role,
	}, nil
}

/* jwt middleware requires config, this function creates that config object */
/*
func NewJWTConfig(secret string) echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTCustomClaims)
		},
		SigningKey: []byte(secret),
	}
}
*/

/* extract the current (logged-in) user from request context */
/*
func CurrentUser(c echo.Context) JWTUser {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaims)

	return JWTUser{
		Id:   claims.Id,
		Role: claims.Role,
	}
}
*/
