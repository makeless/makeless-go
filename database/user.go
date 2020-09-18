package go_saas_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

type User interface {
	GetUser(connection *gorm.DB, user *go_saas_model.User) (*go_saas_model.User, error)
	GetUserByField(connection *gorm.DB, user *go_saas_model.User, field string, value string) (*go_saas_model.User, error)
	CreateUser(connection *gorm.DB, user *go_saas_model.User) (*go_saas_model.User, error)

	UsersTeam(connection *gorm.DB, search string, users []*go_saas_model.User, team *go_saas_model.Team) ([]*go_saas_model.User, error)
}
