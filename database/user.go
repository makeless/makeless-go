package go_saas_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

type User interface {
	GetUser(connection *gorm.DB, user *go_saas_model.User) (*go_saas_model.User, error)
	GetUserByField(connection *gorm.DB, user *go_saas_model.User, field string, value string) (*go_saas_model.User, error)
	CreateUser(connection *gorm.DB, user *go_saas_model.User) (*go_saas_model.User, error)
	IsModelUser(connection *gorm.DB, user *go_saas_model.User, model interface{}) (bool, error)
}
