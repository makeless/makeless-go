package makeless_go_database_basic

import (
	"github.com/jinzhu/gorm"
	"github.com/makeless/makeless-go/model"
)

func (database *Database) UpdatePassword(connection *gorm.DB, user *makeless_go_model.User, newPassword string) (*makeless_go_model.User, error) {
	return user, connection.
		Model(user).
		Update(map[string]interface{}{
			"password": newPassword,
		}).
		Error
}
