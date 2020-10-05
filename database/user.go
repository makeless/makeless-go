package makeless_go_database

import (
	"github.com/makeless/makeless-go/model"
	"github.com/jinzhu/gorm"
)

type User interface {
	GetUser(connection *gorm.DB, user *makeless_go_model.User) (*makeless_go_model.User, error)
	GetUserByField(connection *gorm.DB, user *makeless_go_model.User, field string, value string) (*makeless_go_model.User, error)
	CreateUser(connection *gorm.DB, user *makeless_go_model.User) (*makeless_go_model.User, error)
	IsModelUser(connection *gorm.DB, user *makeless_go_model.User, model interface{}) (bool, error)
}
