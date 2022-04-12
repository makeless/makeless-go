package makeless_go_service_server_password_basic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/makeless/makeless-go/v2/config"
	"github.com/makeless/makeless-go/v2/database/database"
	"github.com/makeless/makeless-go/v2/database/model"
	"github.com/makeless/makeless-go/v2/database/repository"
	"github.com/makeless/makeless-go/v2/mail"
	"github.com/makeless/makeless-go/v2/mailer"
	"github.com/makeless/makeless-go/v2/proto/basic"
	"github.com/makeless/makeless-go/v2/security/auth"
	"github.com/makeless/makeless-go/v2/security/auth_middleware"
	"github.com/makeless/makeless-go/v2/security/crypto"
	"github.com/makeless/makeless-go/v2/security/token"
	"github.com/makeless/makeless-go/v2/service_server/password"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

type PasswordServiceServer struct {
	makeless_go_service_server_password.PasswordServiceServer
	Config              makeless_go_config.Config
	Database            makeless_go_database.Database
	Mailer              makeless_go_mailer.Mailer
	Crypto              makeless_go_crypto.Crypto
	AuthMiddleware      makeless_go_auth_middleware.AuthMiddleware[makeless_go_auth.Claim]
	UserRepository      makeless_go_repository.UserRepository
	GenericRepository   makeless_go_repository.GenericRepository
	PasswordRepository  makeless_go_repository.PasswordRepository
	PasswordRequestMail makeless_go_mail.PasswordRequestMail
	SecurityToken       makeless_go_security_token.SecurityToken
}

func (passwordServiceServer *PasswordServiceServer) UpdatePassword(ctx context.Context, updatePasswordRequest *makeless.UpdatePasswordRequest) (*makeless.UpdatePasswordResponse, error) {
	var err error
	var claim *makeless_go_auth.Claim

	if claim, err = passwordServiceServer.AuthMiddleware.ClaimFromContext(ctx); err != nil {
		return nil, err
	}

	var user = &makeless_go_model.User{
		Model: makeless_go_model.Model{Id: (*claim).GetId()},
		Email: (*claim).GetEmail(),
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

func (passwordServiceServer *PasswordServiceServer) CreatePasswordRequest(ctx context.Context, createPasswordRequestRequest *makeless.CreatePasswordRequestRequest) (*makeless.CreatePasswordRequestResponse, error) {
	var err error
	var userExists bool
	var token string
	var tokenExpire = time.Now().Add(time.Hour * 1)
	var mail makeless_go_mail.Mail

	_, err = passwordServiceServer.UserRepository.GetUserByField(passwordServiceServer.Database.GetConnection().WithContext(ctx), new(makeless_go_model.User), "email", createPasswordRequestRequest.GetEmail())

	if userExists, err = passwordServiceServer.GenericRepository.Exists(err); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if !userExists {
		return nil, status.Errorf(codes.OK, "")
	}

	if token, err = passwordServiceServer.SecurityToken.Generate(32); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var passwordRequest = &makeless_go_model.PasswordRequest{
		Model:  makeless_go_model.Model{Id: uuid.New()},
		Email:  createPasswordRequestRequest.GetEmail(),
		Token:  token,
		Expire: tokenExpire,
		Used:   false,
	}

	if err = passwordServiceServer.PasswordRepository.CreatePasswordRequest(passwordServiceServer.Database.GetConnection().WithContext(ctx), passwordRequest); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if mail, err = passwordServiceServer.PasswordRequestMail.Create(passwordServiceServer.Config, passwordRequest, passwordServiceServer.Config.GetConfiguration().GetLocale()); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = passwordServiceServer.Mailer.SendQueue(mail); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &makeless.CreatePasswordRequestResponse{}, nil
}

func (passwordServiceServer *PasswordServiceServer) ResetPassword(ctx context.Context, resetPasswordRequest *makeless.ResetPasswordRequest) (*makeless.ResetPasswordResponse, error) {
	var err error
	var tx = passwordServiceServer.Database.GetConnection().WithContext(ctx).Begin(new(sql.TxOptions))
	var user *makeless_go_model.User

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err = tx.Error; err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var passwordRequest = &makeless_go_model.PasswordRequest{
		Token: resetPasswordRequest.GetToken(),
	}

	if passwordRequest, err = passwordServiceServer.PasswordRepository.GetPasswordRequest(tx, passwordRequest); err != nil {
		switch errors.Is(err, gorm.ErrRecordNotFound) {
		case true:
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	if user, err = passwordServiceServer.UserRepository.GetUserByField(tx, user, "email", passwordRequest.Email); err != nil {
		switch errors.Is(err, gorm.ErrRecordNotFound) {
		case true:
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	if user.Password, err = passwordServiceServer.Crypto.EncryptPassword(resetPasswordRequest.GetPassword()); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if _, err = passwordServiceServer.PasswordRepository.UpdatePassword(tx, user, user.Password); err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if _, err = passwordServiceServer.PasswordRepository.UpdatePasswordRequest(tx, passwordRequest); err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &makeless.ResetPasswordResponse{}, nil
}
