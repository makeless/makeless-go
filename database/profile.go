package saas_database

import (
	"github.com/loeffel-io/go-saas/model"
)

func (database *Database) UpdateProfile(user *saas_model.User, userId uint) (*saas_model.User, error) {
	return user, database.GetConnection().
		Model(&user).
		Where("users.id = ?", userId).
		Update(map[string]interface{}{
			"name": user.Name,
		}).
		Error
}

func (database *Database) UpdateProfileTeam(team *saas_model.Team, teamId uint, userId uint) (*saas_model.Team, error) {
	return team, database.GetConnection().
		Model(&team).
		Where("teams.id = ? AND teams.user_id = ?", teamId, userId).
		Update(map[string]interface{}{
			"name": team.Name,
		}).
		Error
}
