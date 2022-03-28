package makeless_go_service_server_user

import (
	"context"
	"github.com/makeless/makeless-go/proto/basic"
)

type UserServiceServer interface {
	CreateUser(ctx context.Context, createUserRequest *makeless.CreateUserRequest) (*makeless.CreateUserResponse, error)
	User(ctx context.Context, userRequest *makeless.UserRequest) (*makeless.UserResponse, error)
}
