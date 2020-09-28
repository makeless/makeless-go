package go_saas_security

import (
	"github.com/go-saas/go-saas/database"
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

type Security interface {
	GetDatabase() go_saas_database.Database
	GenerateToken(length int) (string, error)
	UserExists(connection *gorm.DB, field string, value string) (bool, error)
	Login(connection *gorm.DB, field string, value string, password string) (*go_saas_model.User, error)
	Register(connection *gorm.DB, user *go_saas_model.User) (*go_saas_model.User, error)
	EncryptPassword(password string) (string, error)
	ComparePassword(userPassword string, password string) error
	IsTeamUser(connection *gorm.DB, teamId uint, userId uint) (bool, error)
	IsTeamRole(connection *gorm.DB, role string, teamId uint, userId uint) (bool, error)
	IsTeamCreator(connection *gorm.DB, teamId uint, userId uint) (bool, error)
	IsModelUser(connection *gorm.DB, userId uint, model interface{}) (bool, error)
}
