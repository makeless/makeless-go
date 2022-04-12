package makeless_go_service_server_auth_basic

import (
	"context"
	"github.com/makeless/makeless-go/v2/config"
	"github.com/makeless/makeless-go/v2/database/database"
	"github.com/makeless/makeless-go/v2/database/model"
	"github.com/makeless/makeless-go/v2/database/model_transformer"
	"github.com/makeless/makeless-go/v2/database/repository"
	"github.com/makeless/makeless-go/v2/proto/basic"
	"github.com/makeless/makeless-go/v2/security/auth"
	"github.com/makeless/makeless-go/v2/security/auth_middleware"
	"github.com/makeless/makeless-go/v2/security/crypto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type AuthServiceServer struct {
	makeless.AuthServiceServer
	Config            makeless_go_config.Config
	Auth              makeless_go_auth.Auth[makeless_go_auth.Claim]
	Database          makeless_go_database.Database
	Crypto            makeless_go_crypto.Crypto
	AuthMiddleware    makeless_go_auth_middleware.AuthMiddleware[makeless_go_auth.Claim]
	UserRepository    makeless_go_repository.UserRepository
	GenericRepository makeless_go_repository.GenericRepository
	UserTransformer   makeless_go_model_transformer.UserTransformer
}

func (authServiceServer *AuthServiceServer) Login(ctx context.Context, loginRequest *makeless.LoginRequest) (*makeless.LoginResponse, error) {
	var err error
	var token string
	var expireAt time.Time
	var user *makeless_go_model.User

	if user, err = authServiceServer.UserRepository.GetUserByField(authServiceServer.Database.GetConnection().WithContext(ctx), new(makeless_go_model.User), "email", loginRequest.GetEmail()); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if err = authServiceServer.Crypto.ComparePassword(user.Password, loginRequest.GetPassword()); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if token, expireAt, err = authServiceServer.Auth.Sign(user.Id, user.Email); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var authCookie = authServiceServer.Auth.Cookie(token, expireAt)

	if err = grpc.SetHeader(ctx, metadata.Pairs("set-cookie", authCookie.String())); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &makeless.LoginResponse{
		Token:    token,
		ExpireAt: timestamppb.New(expireAt),
	}, nil
}

func (authServiceServer *AuthServiceServer) Logout(ctx context.Context, logoutRequest *makeless.LogoutRequest) (*makeless.LogoutResponse, error) {
	var err error
	var authCookie = authServiceServer.Auth.Cookie("", time.Time{})

	if err = grpc.SetHeader(ctx, metadata.Pairs("set-cookie", authCookie.String())); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &makeless.LogoutResponse{}, nil
}

func (authServiceServer *AuthServiceServer) Refresh(ctx context.Context, refreshRequest *makeless.RefreshRequest) (*makeless.RefreshResponse, error) {
	var err error
	var token string
	var expireAt time.Time
	var claim *makeless_go_auth.Claim

	if claim, err = authServiceServer.AuthMiddleware.ClaimFromContext(ctx); err != nil {
		return nil, err
	}

	var user = &makeless_go_model.User{
		Model: makeless_go_model.Model{Id: (*claim).GetId()},
		Email: (*claim).GetEmail(),
	}

	if token, expireAt, err = authServiceServer.Auth.Sign(user.Id, user.Email); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var authCookie = authServiceServer.Auth.Cookie(token, expireAt)

	if err = grpc.SetHeader(ctx, metadata.Pairs("set-cookie", authCookie.String())); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &makeless.RefreshResponse{
		Token:    token,
		ExpireAt: timestamppb.New(expireAt),
	}, nil
}
