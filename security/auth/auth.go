package makeless_go_auth

import (
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Auth interface {
	Sign(id uuid.UUID, email string) (string, time.Time, error)
	Verify(token string) (*Claim, error)
	Cookie(token string, expireAt time.Time) http.Cookie
}
