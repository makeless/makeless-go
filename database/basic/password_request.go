package go_saas_database_basic

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
	"time"
)

func (database *Database) CreatePasswordRequest(connection *gorm.DB, passwordRequest *go_saas_model.PasswordRequest) error {
	return connection.
		Create(passwordRequest).
		Error
}

func (database *Database) GetPasswordRequest(connection *gorm.DB, passwordRequest *go_saas_model.PasswordRequest) (*go_saas_model.PasswordRequest, error) {
	return passwordRequest, connection.
		Find(passwordRequest, "token = ? AND used = ? AND expire >= ?", passwordRequest.GetToken(), false, time.Now()).
		Error
}

func (database *Database) UpdatePasswordRequest(connection *gorm.DB, passwordRequest *go_saas_model.PasswordRequest) (*go_saas_model.PasswordRequest, error) {
	return passwordRequest, connection.
		Model(passwordRequest).
		Update(map[string]interface{}{
			"used": true,
		}).
		Error
}
