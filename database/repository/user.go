package makeless_go_repository

import (
	"github.com/makeless/makeless-go/v2/database/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(connection *gorm.DB, user *makeless_go_model.User) (*makeless_go_model.User, error)
	GetUser(connection *gorm.DB, user *makeless_go_model.User) (*makeless_go_model.User, error)
	GetUserByField(connection *gorm.DB, user *makeless_go_model.User, field string, value string) (*makeless_go_model.User, error)
	IsModelUser(connection *gorm.DB, user *makeless_go_model.User, model interface{}) (bool, error)
}
