package go_saas_database_basic

import "github.com/go-saas/go-saas/model"

func (database *Database) UpdatePassword(user *go_saas_model.User, newPassword string) (*go_saas_model.User, error) {
	return user, database.GetConnection().
		Model(user).
		Update(map[string]interface{}{
			"password": newPassword,
		}).
		Error
}
