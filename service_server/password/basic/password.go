package makeless_go_password_basic

import (
	"context"
	"github.com/makeless/makeless-go/v2/config"
	"github.com/makeless/makeless-go/v2/database/database"
	"github.com/makeless/makeless-go/v2/database/model"
	"github.com/makeless/makeless-go/v2/database/repository"
	"github.com/makeless/makeless-go/v2/mailer"
	"github.com/makeless/makeless-go/v2/proto/basic"
	"github.com/makeless/makeless-go/v2/security/auth"
	"github.com/makeless/makeless-go/v2/security/auth_middleware"
	"github.com/makeless/makeless-go/v2/security/crypto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PasswordServiceServer struct {
	makeless.PasswordServiceServer
	Config             makeless_go_config.Config
	Database           makeless_go_database.Database
	Mailer             makeless_go_mailer.Mailer
	Crypto             makeless_go_crypto.Crypto
	AuthMiddleware     makeless_go_auth_middleware.AuthMiddleware
	UserRepository     makeless_go_repository.UserRepository
	PasswordRepository makeless_go_repository.PasswordRepository
}

func (passwordServiceServer *PasswordServiceServer) UpdatePassword(ctx context.Context, updatePasswordRequest *makeless.UpdatePasswordRequest) (*makeless.UpdatePasswordResponse, error) {
	var err error
	var claim makeless_go_auth.Claim

	if claim, err = passwordServiceServer.AuthMiddleware.ClaimFromContext(ctx); err != nil {
		return nil, err
	}

	var user = &makeless_go_model.User{
		Model: makeless_go_model.Model{Id: claim.GetId()},
		Email: claim.GetEmail(),
	}

	if user, err = passwordServiceServer.UserRepository.GetUserByField(passwordServiceServer.Database.GetConnection().WithContext(ctx), user, "email", user.Email); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if err = passwordServiceServer.Crypto.ComparePassword(user.Password, updatePasswordRequest.Password); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if user.Password, err = passwordServiceServer.Crypto.EncryptPassword(updatePasswordRequest.GetNewPassword()); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &makeless.UpdatePasswordResponse{}, nil
}
