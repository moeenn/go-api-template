package jwt

import (
	"errors"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

var (
	errInvalidExpired = errors.New("invalid or expired JWT provided")
	errInvalidClaims  = errors.New("failed to parse JWT claims")
	errInvalidScope   = errors.New("JWT scope mismatch")
)

type JWTWithoutExpiry struct {
	Token string `json:"token"`
}

type JWTWithExpiry struct {
	Token  string `json:"token"`
	Expiry int64  `json:"expiry"`
}

type JWTUser struct {
	Id   string `json:"id"`
	Role string `json:"role"`
}

/* create JWT claim and sign using JWT secret */
func NewExpiringJWT(secret string, expiryHours time.Duration, user JWTUser, scope string) (JWTWithExpiry, error) {
	exp := time.Now().UTC().Add(expiryHours)
	expTimestamp := exp.UnixNano() / 1e6 // get JS equivalient timestamp

	claims := &jwtlib.MapClaims{
		"sub":   user.Id,
		"exp":   expTimestamp,
		"role":  user.Role,
		"scope": scope,
	}

	t, err := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return JWTWithExpiry{}, err
	}

	return JWTWithExpiry{
		Token:  t,
		Expiry: expTimestamp,
	}, nil
}

func NewNonExpiringJWT(secret string, user JWTUser, scope string) (JWTWithoutExpiry, error) {
	claims := &jwtlib.MapClaims{
		"sub":   user.Id,
		"role":  user.Role,
		"scope": scope,
	}

	t, err := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return JWTWithoutExpiry{}, err
	}

	return JWTWithoutExpiry{t}, nil
}

/* validate token and extract custom claims */
func ValidateJWT(secret string, tokenString string, isExpiring bool, scope string) (JWTUser, error) {
	var nullUser JWTUser

	parsed, err := jwtlib.Parse(tokenString, func(token *jwtlib.Token) (any, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nullUser, err
	}

	if !parsed.Valid {
		return nullUser, errInvalidExpired
	}

	id, idErr := parsed.Claims.GetSubject()
	if idErr != nil {
		return nullUser, errInvalidClaims
	}

	claims, ok := parsed.Claims.(jwtlib.MapClaims)
	if !ok {
		return nullUser, errInvalidClaims
	}

	roleRaw, roleRawErr := claims["role"]
	scopeRaw, scopeRawErr := claims["scope"]

	if !roleRawErr || !scopeRawErr {
		return nullUser, errInvalidClaims
	}

	role, roleOk := roleRaw.(string)
	gotScope, gotScopeOk := scopeRaw.(string)

	if !roleOk || !gotScopeOk {
		return nullUser, errInvalidClaims
	}

	if gotScope != scope {
		return nullUser, errInvalidScope
	}

	if isExpiring {
		expiry, expErr := parsed.Claims.GetExpirationTime()
		if expErr != nil {
			return nullUser, errInvalidClaims
		}

		if expiry.Before(time.Now()) {
			return nullUser, errInvalidExpired
		}
	}

	return JWTUser{
		Id:   id,
		Role: role,
	}, nil
}
