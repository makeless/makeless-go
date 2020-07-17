package go_saas_basic_database

import (
	"github.com/go-saas/go-saas/model"
	"golang.org/x/sync/errgroup"
)

func (database *Database) CreateTeam(team *go_saas_model.Team, teamUser *go_saas_model.TeamUser) (*go_saas_model.Team, error) {
	var g = new(errgroup.Group)

	g.Go(func() error {
		return database.GetConnection().Set("gorm:save_associations", false).Create(teamUser).Error
	})

	g.Go(func() error {
		return database.GetConnection().Set("gorm:save_associations", false).Create(team).Error
	})

	return team, g.Wait()
}

func (database *Database) DeleteTeamUsers(user *go_saas_model.User, team *go_saas_model.Team) error {
	if user.GetId() != *team.GetUserId() {
		return database.GetConnection().
			Raw("DELETE FROM team_users WHERE team_users.team_id = ? AND team_users.user_id = ?", team.GetId(), user.GetId()).
			Error
	}

	return database.GetConnection().
		Raw("DELETE FROM team_users WHERE team_users.team_id = ?", team.GetId()).
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
		Raw("SELECT COUNT(*) FROM team_users WHERE team_users.team_id = ? AND team_users.user_id = ? LIMIT 1", team.GetId(), user.GetId()).
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
