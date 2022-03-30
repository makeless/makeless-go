package makeless_go_auth_middleware

import (
	"context"
)

type AuthMiddleware interface {
	AuthFunc(ctx context.Context) (context.Context, error)
	TokenLookup(ctx context.Context) (string, bool, error)
}
