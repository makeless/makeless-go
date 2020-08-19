package go_saas_database_basic

import "github.com/go-saas/go-saas/model"

func (database *Database) UpdateProfileTeam(team *go_saas_model.Team) (*go_saas_model.Team, error) {
	return team, database.GetConnection().
		Model(&team).
		Update(map[string]interface{}{
			"name": team.Name,
		}).
		Error
}
