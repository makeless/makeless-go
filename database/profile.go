package saas_database

import (
	"github.com/go-saas/go-saas/model"
)

func (database *Database) UpdateProfile(user *go_saas_model.User, userId uint) (*go_saas_model.User, error) {
	return user, database.GetConnection().
		Model(&user).
		Where("users.id = ?", userId).
		Update(map[string]interface{}{
			"name": user.Name,
		}).
		Error
}

func (database *Database) UpdateProfileTeam(team *go_saas_model.Team, teamId uint, userId uint) (*go_saas_model.Team, error) {
	return team, database.GetConnection().
		Model(&team).
		Where("teams.id = ? AND teams.user_id = ?", teamId, userId).
		Update(map[string]interface{}{
			"name": team.Name,
		}).
		Error
}
