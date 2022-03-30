package makeless_go_auth_basic

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Auth struct {
	Key               string
	KeySigningMethod  jwt.SigningMethod
	KeyExpireDuration time.Duration
	CookieDomain      string
	CookieSameSite    http.SameSite
}

func (auth *Auth) Sign(id uuid.UUID, email string) (string, time.Time, error) {
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

func (auth *Auth) Verify(token string) (*Claim, error) {
	var err error
	var ok bool
	var claim = new(Claim)
	var jwtToken *jwt.Token

	if jwtToken, err = jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.Key), nil
	}); err != nil {
		return nil, fmt.Errorf("unable to parse token: %s", err.Error())
	}

	if claim, ok = jwtToken.Claims.(*Claim); !ok || !jwtToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claim, nil
}

func (auth *Auth) Cookie(token string, expireAt time.Time) http.Cookie {
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
