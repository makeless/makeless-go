package makeless_go_service_server_auth

import (
	"context"
	"github.com/makeless/makeless-go/v2/proto/basic"
)

type AuthServiceServer interface {
	Login(ctx context.Context, loginRequest *makeless.LoginRequest) (*makeless.LoginResponse, error)
	Logout(ctx context.Context, logoutRequest *makeless.LogoutRequest) (*makeless.LogoutResponse, error)
	Refresh(ctx context.Context, refreshRequest *makeless.RefreshRequest) (*makeless.RefreshResponse, error)
}
