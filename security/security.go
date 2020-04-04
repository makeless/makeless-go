package saas_security

import (
	"github.com/loeffel-io/go-saas/model"
)

type Security interface {
	Login(username string, password string) (*saas_model.User, error)
	Register(user *saas_model.User) (*saas_model.User, error)
}
