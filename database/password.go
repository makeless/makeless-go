package saas_database

import "github.com/go-saas/go-saas/model"

func (database *Database) UpdatePassword(userId uint, password string) error {
	return database.GetConnection().
		Model(&go_saas_model.User{
			Model: go_saas_model.Model{Id: userId},
		}).
		Where("users.id = ?", userId).
		Update(map[string]interface{}{
			"password": password,
		}).
		Error

}
