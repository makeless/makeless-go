package go_saas_basic_database

import (
	"github.com/go-saas/go-saas/model"
)

func (database *Database) CreateTeam(user *go_saas_model.User, team *go_saas_model.Team) (*go_saas_model.Team, error) {
	return team, database.GetConnection().
		Model(user).
		Association("Teams").
		Append(team).
		Error
}

func (database *Database) LeaveTeam(user *go_saas_model.User, team *go_saas_model.Team) error {
	return database.GetConnection().
		Model(user).
		Association("Teams").
		Delete(team).
		Error
}

func (database *Database) DeleteTeam(team *go_saas_model.Team) error {
	return database.GetConnection().
		Unscoped().
		Where("teams.id = ? AND teams.user_id = ?", team.GetId(), team.GetUserId()).
		Delete(team).
		Error
}

func (database *Database) IsTeamMember(team *go_saas_model.Team, user *go_saas_model.User) (bool, error) {
	var count int

	return count == 1, database.GetConnection().
		Raw("SELECT COUNT(*) FROM user_teams WHERE user_teams.team_id = ? AND user_teams.user_id = ? LIMIT 1", team.GetId(), user.GetId()).
		Count(&count).
		Error
}

func (database *Database) IsTeamOwner(team *go_saas_model.Team, user *go_saas_model.User) (bool, error) {
	var count int

	return count == 1, database.GetConnection().
		Raw("SELECT COUNT(*) FROM teams WHERE teams.id = ? AND teams.user_id = ? LIMIT 1", team.GetId(), user.GetId()).
		Count(&count).
		Error
}
