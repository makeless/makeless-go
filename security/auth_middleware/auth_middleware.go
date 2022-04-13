package makeless_go_auth_middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddleware[T jwt.Claims] interface {
	AuthFunc(ctx context.Context) (context.Context, error)
	TokenFromContext(ctx context.Context) (string, bool, error)
	ClaimFromContext(ctx context.Context) (T, error)
}
