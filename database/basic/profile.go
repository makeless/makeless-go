package go_saas_basic_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/struct"
)

func (database *Database) UpdateProfile(user *go_saas_model.User, profile *_struct.Profile) (*go_saas_model.User, error) {
	return user, database.GetConnection().
		Model(&user).
		Update(map[string]interface{}{
			"name": profile.GetName(),
		}).
		Error
}
