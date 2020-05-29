package saas_security

import (
	"github.com/go-saas/go-saas/model"
)

type Security interface {
	Login(field string, id string, password string) (*saas_model.User, error)
	Register(user *saas_model.User) (*saas_model.User, error)
	TokenLogin(token string) (*saas_model.User, error)
	EncryptPassword(password string) ([]byte, error)

	IsTeamMember(teamId uint, userId uint) (bool, error)
	IsTeamOwner(teamId uint, userId uint) (bool, error)
}
