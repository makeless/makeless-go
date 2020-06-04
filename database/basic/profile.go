package go_saas_basic_database

import (
	"github.com/go-saas/go-saas/model"
)

func (database *Database) UpdateProfile(user *go_saas_model.User) (*go_saas_model.User, error) {
	return user, database.GetConnection().
		Model(&user).
		Update(map[string]interface{}{
			"name": user.Name,
		}).
		Error
}

func (database *Database) UpdateProfileTeam(team *go_saas_model.Team) (*go_saas_model.Team, error) {
	return team, database.GetConnection().
		Model(&team).
		Update(map[string]interface{}{
			"name": team.Name,
		}).
		Error
}
