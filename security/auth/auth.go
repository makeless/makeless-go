package makeless_go_auth

import (
	"github.com/google/uuid"
	"time"
)

type Auth interface {
	Sign(id uuid.UUID, email string) (string, time.Time, error)
}
