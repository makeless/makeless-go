package go_saas_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

type PasswordRequest interface {
	CreatePasswordRequest(connection *gorm.DB, passwordRequest *go_saas_model.PasswordRequest) error
	GetPasswordRequest(connection *gorm.DB, passwordRequest *go_saas_model.PasswordRequest) (*go_saas_model.PasswordRequest, error)
	UpdatePasswordRequest(connection *gorm.DB, passwordRequest *go_saas_model.PasswordRequest) (*go_saas_model.PasswordRequest, error)
}
