package makeless_go_auth_basic

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

type Auth struct {
	Key            string
	SigningMethod  jwt.SigningMethod
	ExpireDuration time.Duration
}

func (auth *Auth) Sign(id uuid.UUID, email string) (string, time.Time, error) {
	var err error
	var token string
	var expireAt = time.Now().UTC().Add(auth.ExpireDuration)
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
