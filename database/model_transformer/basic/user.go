package makeless_go_model_transformer_basic

import (
	"github.com/makeless/makeless-go/v2/database/model"
	"github.com/makeless/makeless-go/v2/proto/basic"
	"strings"
)

type UserTransformer struct {
}

func (userTransformer *UserTransformer) CreateUserRequestToUser(createUserRequest *makeless.CreateUserRequest, token string) (*makeless_go_model.User, error) {
	return &makeless_go_model.User{
		Name:     createUserRequest.GetName(),
		Email:    strings.ToLower(createUserRequest.GetEmail()),
		Password: createUserRequest.GetPassword(),
		EmailVerification: &makeless_go_model.EmailVerification{
			Token:    token,
			Verified: false,
		},
	}, nil
}

// FIXME: Add email verification
func (userTransformer *UserTransformer) UserToUser(user *makeless_go_model.User) (*makeless.User, error) {
	return &makeless.User{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
