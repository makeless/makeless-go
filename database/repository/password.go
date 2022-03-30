package makeless_go_repository

import (
	"github.com/makeless/makeless-go/v2/database/model"
	"gorm.io/gorm"
)

type PasswordRepository interface {
	UpdatePassword(connection *gorm.DB, user *makeless_go_model.User, newPassword string) (*makeless_go_model.User, error)
	CreatePasswordRequest(connection *gorm.DB, passwordRequest *makeless_go_model.PasswordRequest) error
	GetPasswordRequest(connection *gorm.DB, passwordRequest *makeless_go_model.PasswordRequest) (*makeless_go_model.PasswordRequest, error)
	UpdatePasswordRequest(connection *gorm.DB, passwordRequest *makeless_go_model.PasswordRequest) (*makeless_go_model.PasswordRequest, error)
}
