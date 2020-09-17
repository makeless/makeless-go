package go_saas_database_basic

import (
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/struct"
	"github.com/jinzhu/gorm"
)

func (database *Database) UpdateProfile(connection *gorm.DB, user *go_saas_model.User, profile *_struct.Profile) (*go_saas_model.User, error) {
	return user, connection.
		Model(&user).
		Update(map[string]interface{}{
			"name": profile.GetName(),
		}).
		Error
}
