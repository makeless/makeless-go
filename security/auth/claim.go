package makeless_go_auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Claim interface {
	jwt.Claims
	GetId() uuid.UUID
	GetEmail() string
}
