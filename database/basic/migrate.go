package go_saas_database_basic

import "github.com/go-saas/go-saas/model"

func (database *Database) Migrate() error {
	return database.GetConnection().AutoMigrate(
		new(go_saas_model.User),
		new(go_saas_model.Token),
		new(go_saas_model.Team),
		new(go_saas_model.TeamUser),
		new(go_saas_model.TeamInvitation),
		new(go_saas_model.PasswordRequest),
	).Error
}
