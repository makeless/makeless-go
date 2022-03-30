package makeless_go_service_server_user

import (
	"context"
	"github.com/makeless/makeless-go/v2/proto/basic"
)

type UserServiceServer interface {
	CreateUser(ctx context.Context, createUserRequest *makeless.CreateUserRequest) (*makeless.CreateUserResponse, error)
	CurrentUser(ctx context.Context, userRequest *makeless.CurrentUserRequest) (*makeless.CurrentUserResponse, error)
}
