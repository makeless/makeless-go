package go_saas_basic_database

import "github.com/go-saas/go-saas/model"

func (database *Database) Migrate() error {
	return database.GetConnection().AutoMigrate(
		new(go_saas_model.User),
		new(go_saas_model.Token),
		new(go_saas_model.Team),
	).Error
}
