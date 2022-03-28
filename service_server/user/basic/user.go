package makeless_go_service_server_user_basic

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/makeless/makeless-go/config"
	"github.com/makeless/makeless-go/database/database"
	"github.com/makeless/makeless-go/database/model"
	"github.com/makeless/makeless-go/database/model_transformer"
	"github.com/makeless/makeless-go/database/repository"
	"github.com/makeless/makeless-go/mail"
	"github.com/makeless/makeless-go/mailer"
	"github.com/makeless/makeless-go/proto/basic"
	"github.com/makeless/makeless-go/security/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserServiceServer struct {
	makeless.UserServiceServer
	Config                makeless_go_config.Config
	Database              makeless_go_database.Database
	Mailer                makeless_go_mailer.Mailer
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

	if !userExists {
		return nil, status.Errorf(codes.AlreadyExists, "%s", user.Email)
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

func (userServiceServer *UserServiceServer) User(ctx context.Context, userRequest *makeless.UserRequest) (*makeless.UserResponse, error) {
	var err error
	var user = &makeless_go_model.User{
		Model: makeless_go_model.Model{Id: uuid.New()},
	}

	if user, err = userServiceServer.UserRepository.GetUser(userServiceServer.Database.GetConnection().WithContext(ctx), user); err != nil {
		switch errors.Is(err, gorm.ErrRecordNotFound) {
		case true:
			return nil, status.Errorf(codes.NotFound, err.Error())
		case false:
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	var userResponseUser *makeless.User
	if userResponseUser, err = userServiceServer.UserTransformer.UserToUser(user); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &makeless.UserResponse{
		User: userResponseUser,
	}, nil
}
