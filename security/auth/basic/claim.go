package makeless_go_auth_basic

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Claim struct {
	jwt.RegisteredClaims
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

func (claim *Claim) GetId() uuid.UUID {
	return claim.Id
}

func (claim *Claim) GetEmail() string {
	return claim.Email
}
