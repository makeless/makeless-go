package makeless_go_password

import (
	"context"
	"github.com/makeless/makeless-go/v2/proto/basic"
)

type PasswordServiceServer interface {
	UpdatePassword(ctx context.Context, updatePasswordRequest *makeless.UpdatePasswordRequest) (*makeless.UpdatePasswordResponse, error)
	CreatePasswordRequest(ctx context.Context, createPasswordRequestRequest *makeless.CreatePasswordRequestRequest) (*makeless.CreatePasswordRequestResponse, error)
	ResetPassword(ctx context.Context, resetPasswordRequest *makeless.ResetPasswordRequest) (*makeless.ResetPasswordResponse, error)
}
