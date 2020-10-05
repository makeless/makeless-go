package makeless_go_database_basic

import (
	"gorm.io/gorm"
	"github.com/makeless/makeless-go/model"
	"time"
)

func (database *Database) CreatePasswordRequest(connection *gorm.DB, passwordRequest *makeless_go_model.PasswordRequest) error {
	return connection.
		Create(passwordRequest).
		Error
}

func (database *Database) GetPasswordRequest(connection *gorm.DB, passwordRequest *makeless_go_model.PasswordRequest) (*makeless_go_model.PasswordRequest, error) {
	return passwordRequest, connection.
		Find(passwordRequest, "token = ? AND used = ? AND expire >= ?", passwordRequest.GetToken(), false, time.Now()).
		Error
}

func (database *Database) UpdatePasswordRequest(connection *gorm.DB, passwordRequest *makeless_go_model.PasswordRequest) (*makeless_go_model.PasswordRequest, error) {
	return passwordRequest, connection.
		Model(passwordRequest).
		Update(map[string]interface{}{
			"used": true,
		}).
		Error
}
