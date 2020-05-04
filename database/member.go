package saas_database

import (
	"github.com/loeffel-io/go-saas/model"
)

func (database *Database) GetMembers(users []*saas_model.User, teamId uint, userId uint) ([]*saas_model.User, error) {
	return users, database.GetConnection().
		Select("users.id, users.name").
		Preload("Teams", "teams.id = ? AND teams.user_id = ?", teamId, userId).
		Find(&users).
		Error
}
