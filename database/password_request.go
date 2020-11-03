package makeless_go_database

import (
	"github.com/makeless/makeless-go/model"
	"gorm.io/gorm"
)

type PasswordRequest interface {
	CreatePasswordRequest(connection *gorm.DB, passwordRequest *makeless_go_model.PasswordRequest) error
	GetPasswordRequest(connection *gorm.DB, passwordRequest *makeless_go_model.PasswordRequest) (*makeless_go_model.PasswordRequest, error)
	UpdatePasswordRequest(connection *gorm.DB, passwordRequest *makeless_go_model.PasswordRequest) (*makeless_go_model.PasswordRequest, error)
}
