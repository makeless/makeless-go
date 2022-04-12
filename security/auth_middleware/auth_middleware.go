package makeless_go_auth_middleware

import (
	"context"
)

type AuthMiddleware[T any] interface {
	AuthFunc(ctx context.Context) (context.Context, error)
	TokenLookup(ctx context.Context) (string, bool, error)
	ClaimFromContext(ctx context.Context) (*T, error)
}
