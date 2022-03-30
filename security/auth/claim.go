package makeless_go_auth

import "github.com/google/uuid"

type Claim interface {
	GetId() uuid.UUID
	GetEmail() string
}
