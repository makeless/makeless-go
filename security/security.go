package makeless_go_security

import (
	"github.com/makeless/makeless-go/database"
	"github.com/makeless/makeless-go/model"
	"gorm.io/gorm"
)

type Security interface {
	GetDatabase() makeless_go_database.Database
	GenerateToken(length int) (string, error)
	UserExists(connection *gorm.DB, field string, value string) (bool, error)
	Login(connection *gorm.DB, field string, value string, password string) (*makeless_go_model.User, error)
	Register(connection *gorm.DB, user *makeless_go_model.User) (*makeless_go_model.User, error)
	EncryptPassword(password string) (string, error)
	ComparePassword(userPassword string, password string) error
	IsTeamUser(connection *gorm.DB, teamId uint, userId uint) (bool, error)
	IsTeamRole(connection *gorm.DB, role string, teamId uint, userId uint) (bool, error)
	IsTeamCreator(connection *gorm.DB, teamId uint, userId uint) (bool, error)
	IsNotTeamCreator(connection *gorm.DB, teamId uint, userId uint) (bool, error)
	IsModelUser(connection *gorm.DB, userId uint, model interface{}) (bool, error)
	IsModelTeam(connection *gorm.DB, teamId uint, model interface{}) (bool, error)
}
