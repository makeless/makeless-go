package go_saas_security

import (
	"github.com/go-saas/go-saas/database"
	"github.com/go-saas/go-saas/model"
)

type Security interface {
	GetDatabase() go_saas_database.Database
	Login(field string, value string, password string) (*go_saas_model.User, error)
	Register(user *go_saas_model.User) (*go_saas_model.User, error)
	TokenLogin(value string) (*go_saas_model.User, *go_saas_model.Team, error)
	EncryptPassword(password string) (string, error)
	ComparePassword(userPassword string, password string) error
	IsTeamMember(teamId uint, userId uint) (bool, error)
	IsTeamOwner(teamId uint, userId uint) (bool, error)
	IsTeamCreator(teamId uint, userId uint) (bool, error)
}
