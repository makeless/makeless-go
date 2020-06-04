package go_saas_security

import (
	"github.com/go-saas/go-saas/model"
)

type Security interface {
	Login(field string, id string, password string) (*go_saas_model.User, error)
	Register(user *go_saas_model.User) (*go_saas_model.User, error)
	TokenLogin(token string) (*go_saas_model.User, error)

	EncryptPassword(password string) ([]byte, error)
	ComparePassword(userPassword string, password string) error

	IsTeamMember(teamId uint, userId uint) (bool, error)
	IsTeamOwner(teamId uint, userId uint) (bool, error)
}
