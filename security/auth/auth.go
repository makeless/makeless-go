package makeless_go_auth

import (
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type Auth[T jwt.Claims] interface {
	GetKeyExpireDuration() time.Duration
	GetCookieDomain() string
	Sign(claim T) (string, error)
	Verify(token string) (T, error)
	Cookie(token string, expireAt time.Time) http.Cookie
}
