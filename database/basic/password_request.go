package go_saas_database_basic

import (
	"github.com/go-saas/go-saas/model"
	"time"
)

func (database *Database) CreatePasswordRequest(passwordRequest *go_saas_model.PasswordRequest) error {
	return database.GetConnection().
		Create(passwordRequest).
		Error
}

func (database *Database) GetPasswordRequest(passwordRequest *go_saas_model.PasswordRequest) (*go_saas_model.PasswordRequest, error) {
	return passwordRequest, database.GetConnection().
		Find(passwordRequest, "token = ? AND used = ? AND expire >= ?", passwordRequest.GetToken(), false, time.Now()).
		Error
}

func (database *Database) UpdatePasswordRequest(passwordRequest *go_saas_model.PasswordRequest) (*go_saas_model.PasswordRequest, error) {
	return passwordRequest, database.GetConnection().
		Model(passwordRequest).
		Update(map[string]interface{}{
			"used": true,
		}).
		Error
}
