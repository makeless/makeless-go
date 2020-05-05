package saas_database

import "github.com/loeffel-io/go-saas/model"

func (database *Database) UpdatePassword(userId uint, password string) error {
	return database.GetConnection().
		Model(&saas_model.User{
			Model: saas_model.Model{Id: userId},
		}).
		Where("users.id = ?", userId).
		Update(map[string]interface{}{
			"password": password,
		}).
		Error

}
