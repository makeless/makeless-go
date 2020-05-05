package saas_security

import (
	"github.com/loeffel-io/go-saas/model"
)

type Security interface {
	Login(field string, id string, password string) (*saas_model.User, error)
	Register(user *saas_model.User) (*saas_model.User, error)
	TokenLogin(token string) (*saas_model.User, error)
	EncryptPassword(password string) ([]byte, error)
}
