package go_saas_database_basic

import (
	"fmt"
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

// CreateTeam creates team
func (database *Database) CreateTeam(connection *gorm.DB, team *go_saas_model.Team) (*go_saas_model.Team, error) {
	return team, connection.
		Create(team).
		Error
}

// CreateTeam retrieves team
func (database *Database) GetTeam(connection *gorm.DB, team *go_saas_model.Team) (*go_saas_model.Team, error) {
	return team, connection.
		Preload("TeamUsers.Team").
		Preload("TeamUsers.User").
		Preload("TeamInvitations").
		Find(team).
		Error
}

// AddTeamInvitations appends team invitations to a given team
func (database *Database) AddTeamInvitations(connection *gorm.DB, team *go_saas_model.Team, teamInvitations []*go_saas_model.TeamInvitation) (*go_saas_model.Team, error) {
	return team, connection.
		Model(team).
		Association("TeamInvitations").
		Append(teamInvitations).
		Error
}

// GetTeamUsers retrieves team users by search
func (database *Database) GetTeamUsers(connection *gorm.DB, search string, teamUsers []*go_saas_model.TeamUser, team *go_saas_model.Team) ([]*go_saas_model.TeamUser, error) {
	var query = connection

	if search != "" {
		query = query.Where(
			"users.name LIKE ? OR users.email LIKE ?",
			fmt.Sprintf(`%s%s%s`, "%", search, "%"),
			fmt.Sprintf(`%s%s%s`, "%", search, "%"),
		)
	}

	return teamUsers, query.
		Joins("JOIN users ON team_users.user_id = users.id").
		Where("team_users.team_id = ", team.GetId()).
		Find(&teamUsers).
		Error
}

// AddTeamUser adds teamUsers to team
func (database *Database) AddTeamUsers(connection *gorm.DB, teamUsers []*go_saas_model.TeamUser, team *go_saas_model.Team) error {
	return connection.
		Model(team).
		Association("TeamUsers").
		Append(teamUsers).
		Error
}

// DeleteTeamUser deletes teamUser
func (database *Database) DeleteTeamUser(connection *gorm.DB, user *go_saas_model.User, team *go_saas_model.Team) error {
	return connection.
		Exec("DELETE FROM team_users WHERE team_users.team_id = ? AND team_users.user_id = ?", team.GetId(), user.GetId()).
		Error
}

// DeleteTeam deletes team and all their teamUsers
func (database *Database) DeleteTeam(connection *gorm.DB, team *go_saas_model.Team) error {
	return connection.
		Select("TeamUsers", "TeamInvitations").
		Delete(team).
		Error
}

// IsTeamUser checks if user is part of team
func (database *Database) IsTeamUser(connection *gorm.DB, team *go_saas_model.Team, user *go_saas_model.User) (bool, error) {
	var count int

	return count == 1, connection.
		Raw("SELECT COUNT(*) FROM team_users WHERE team_users.team_id = ? AND team_users.user_id = ? LIMIT 1", team.GetId(), user.GetId()).
		Count(&count).
		Error
}

// IsTeamRole checks if user is part of team and has given role
func (database *Database) IsTeamRole(connection *gorm.DB, role string, team *go_saas_model.Team, user *go_saas_model.User) (bool, error) {
	var count int

	return count == 1, connection.
		Raw("SELECT COUNT(*) FROM team_users WHERE team_users.team_id = ? AND team_users.user_id = ? AND team_users.role = ? LIMIT 1", team.GetId(), user.GetId(), role).
		Count(&count).
		Error
}

// IsTeamCreator checks if user is team creator
func (database *Database) IsTeamCreator(connection *gorm.DB, team *go_saas_model.Team, user *go_saas_model.User) (bool, error) {
	var count int

	return count == 1, connection.
		Raw("SELECT COUNT(*) FROM teams WHERE teams.id = ? AND teams.user_id = ? LIMIT 1", team.GetId(), user.GetId()).
		Count(&count).
		Error
}

// IsNotTeamCreator checks if user is not team creator
func (database *Database) IsNotTeamCreator(connection *gorm.DB, team *go_saas_model.Team, user *go_saas_model.User) (bool, error) {
	var count int

	return count == 0, connection.
		Raw("SELECT COUNT(*) FROM teams WHERE teams.id = ? AND teams.user_id = ? LIMIT 1", team.GetId(), user.GetId()).
		Count(&count).
		Error
}
