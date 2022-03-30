package makeless_go_repository_basic

import (
	"github.com/makeless/makeless-go/v2/database/model"
	"gorm.io/gorm"
	"time"
)

type PasswordRepository struct {
}

func (passwordRepository *PasswordRepository) UpdatePassword(connection *gorm.DB, user *makeless_go_model.User, newPassword string) (*makeless_go_model.User, error) {
	return user, connection.
		Model(user).
		Updates(map[string]interface{}{
			"password": newPassword,
		}).
		Error
}

func (passwordRepository *PasswordRepository) CreatePasswordRequest(connection *gorm.DB, passwordRequest *makeless_go_model.PasswordRequest) error {
	return connection.
		Create(passwordRequest).
		Error
}

func (passwordRepository *PasswordRepository) GetPasswordRequest(connection *gorm.DB, passwordRequest *makeless_go_model.PasswordRequest) (*makeless_go_model.PasswordRequest, error) {
	return passwordRequest, connection.
		First(passwordRequest, "token = ? AND used = ? AND expire >= ?", passwordRequest.Token, false, time.Now()).
		Error
}

func (passwordRepository *PasswordRepository) UpdatePasswordRequest(connection *gorm.DB, passwordRequest *makeless_go_model.PasswordRequest) (*makeless_go_model.PasswordRequest, error) {
	return passwordRequest, connection.
		Model(passwordRequest).
		Updates(map[string]interface{}{
			"used": true,
		}).
		Error
}
