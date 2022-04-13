package makeless_go_auth_basic

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Auth[T jwt.Claims] struct {
	Claim             T
	Key               string
	KeySigningMethod  jwt.SigningMethod
	KeyExpireDuration time.Duration
	CookieDomain      string
	CookieSameSite    http.SameSite
}

func (auth *Auth[T]) Sign(id uuid.UUID, email string) (string, time.Time, error) {
	var err error
	var token string
	var expireAt = time.Now().UTC().Add(auth.KeyExpireDuration)
	var issuedAt = time.Now().UTC()

	var claim = Claim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
			IssuedAt:  jwt.NewNumericDate(issuedAt),
		},
		Id:    id,
		Email: email,
	}

	if token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(auth.Key)); err != nil {
		return "", expireAt, err
	}

	return token, expireAt, nil
}

func (auth *Auth[T]) Verify(token string) (T, error) {
	var err error
	var ok bool
	var jwtToken *jwt.Token
	var claim = auth.Claim

	if jwtToken, err = jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.Key), nil
	}); err != nil {
		return claim, fmt.Errorf("unable to parse token: %s", err.Error())
	}

	if claim, ok = jwtToken.Claims.(T); !ok || !jwtToken.Valid {
		return claim, fmt.Errorf("invalid token")
	}

	return claim, nil
}

func (auth *Auth[T]) Cookie(token string, expireAt time.Time) http.Cookie {
	return http.Cookie{
		Name:     "jwt",
		Value:    token,
		Domain:   auth.CookieDomain,
		Expires:  expireAt,
		Secure:   true,
		HttpOnly: true,
		SameSite: auth.CookieSameSite,
	}
}
