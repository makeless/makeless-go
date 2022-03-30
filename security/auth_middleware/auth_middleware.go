package makeless_go_auth_middleware

import (
	"context"
	"github.com/makeless/makeless-go/v2/security/auth"
)

type AuthMiddleware interface {
	AuthFunc(ctx context.Context) (context.Context, error)
	TokenLookup(ctx context.Context) (string, bool, error)
	ClaimFromContext(ctx context.Context) (makeless_go_auth.Claim, error)
}
