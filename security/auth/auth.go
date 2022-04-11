package makeless_go_auth

import (
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Auth[T any] interface {
	Sign(id uuid.UUID, email string) (string, time.Time, error)
	Verify[T](token string) (T, error)
	Cookie(token string, expireAt time.Time) http.Cookie
}
