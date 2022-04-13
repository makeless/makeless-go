package makeless_go_service_server_user_basic

import (
	"context"
	"errors"
	"github.com/makeless/makeless-go/v2/config"
	"github.com/makeless/makeless-go/v2/database/database"
	"github.com/makeless/makeless-go/v2/database/model"
	"github.com/makeless/makeless-go/v2/database/model_transformer"
	"github.com/makeless/makeless-go/v2/database/repository"
	"github.com/makeless/makeless-go/v2/mail"
	"github.com/makeless/makeless-go/v2/mailer"
	"github.com/makeless/makeless-go/v2/proto/basic"
	"github.com/makeless/makeless-go/v2/security/auth/basic"
	"github.com/makeless/makeless-go/v2/security/auth_middleware"
	"github.com/makeless/makeless-go/v2/security/crypto"
	"github.com/makeless/makeless-go/v2/security/token"
	"github.com/makeless/makeless-go/v2/service_server/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserServiceServer struct {
	makeless_go_service_server_user.UserServiceServer
	Config                makeless_go_config.Config
	Database              makeless_go_database.Database
	Mailer                makeless_go_mailer.Mailer
	Crypto                makeless_go_crypto.Crypto
	AuthMiddleware        makeless_go_auth_middleware.AuthMiddleware[*makeless_go_auth_basic.Claim]
	UserRepository        makeless_go_repository.UserRepository
	GenericRepository     makeless_go_repository.GenericRepository
	UserTransformer       makeless_go_model_transformer.UserTransformer
	SecurityToken         makeless_go_security_token.SecurityToken
	EmailVerificationMail makeless_go_mail.EmailVerificationMail
}

func (userServiceServer *UserServiceServer) CreateUser(ctx context.Context, createUserRequest *makeless.CreateUserRequest) (*makeless.CreateUserResponse, error) {
	var err error
	var token string
	var mail makeless_go_mail.Mail
	var user *makeless_go_model.User
	var userExists bool

	if token, err = userServiceServer.SecurityToken.Generate(32); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if user, err = userServiceServer.UserTransformer.CreateUserRequestToUser(createUserRequest, token); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	_, err = userServiceServer.UserRepository.GetUserByField(userServiceServer.Database.GetConnection().WithContext(ctx), new(makeless_go_model.User), "email", user.Email)

	if userExists, err = userServiceServer.GenericRepository.Exists(err); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if userExists {
		return nil, status.Errorf(codes.AlreadyExists, "%s", user.Email)
	}

	if user.Password, err = userServiceServer.Crypto.EncryptPassword(user.Password); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if user, err = userServiceServer.UserRepository.CreateUser(userServiceServer.Database.GetConnection().WithContext(ctx), user); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if mail, err = userServiceServer.EmailVerificationMail.Create(userServiceServer.Config, user, userServiceServer.Config.GetConfiguration().GetLocale()); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = userServiceServer.Mailer.SendQueue(mail); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var createUserResponseUser *makeless.User
	if createUserResponseUser, err = userServiceServer.UserTransformer.UserToUser(user); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &makeless.CreateUserResponse{
		User: createUserResponseUser,
	}, nil
}

func (userServiceServer UserServiceServer) CurrentUser(ctx context.Context, currentUserRequest *makeless.CurrentUserRequest) (*makeless.CurrentUserResponse, error) {
	var err error
	var claim *makeless_go_auth_basic.Claim

	if claim, err = userServiceServer.AuthMiddleware.ClaimFromContext(ctx); err != nil {
		return nil, err
	}

	var user = &makeless_go_model.User{
		Model: makeless_go_model.Model{Id: (*claim).GetId()},
	}

	if user, err = userServiceServer.UserRepository.GetUser(userServiceServer.Database.GetConnection().WithContext(ctx), user); err != nil {
		switch errors.Is(err, gorm.ErrRecordNotFound) {
		case true:
			return nil, status.Errorf(codes.NotFound, err.Error())
		case false:
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	var currentUserResponseUser *makeless.User
	if currentUserResponseUser, err = userServiceServer.UserTransformer.UserToUser(user); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &makeless.CurrentUserResponse{
		User: currentUserResponseUser,
	}, nil
}
