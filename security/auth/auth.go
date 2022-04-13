package makeless_go_auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Auth[T jwt.Claims] interface {
	Sign(id uuid.UUID, email string) (string, time.Time, error)
	Verify(token string) (*T, error)
	Cookie(token string, expireAt time.Time) http.Cookie
}
