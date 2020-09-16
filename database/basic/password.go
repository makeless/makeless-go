package go_saas_database_basic

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

func (database *Database) UpdatePassword(connection *gorm.DB, user *go_saas_model.User, newPassword string) (*go_saas_model.User, error) {
	return user, connection.
		Model(user).
		Update(map[string]interface{}{
			"password": newPassword,
		}).
		Error
}
