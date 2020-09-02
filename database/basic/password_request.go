package go_saas_database_basic

import "github.com/go-saas/go-saas/model"

func (database *Database) CreatePasswordRequest(passwordRequest *go_saas_model.PasswordRequest) error {
	return database.GetConnection().
		Create(passwordRequest).
		Error
}
