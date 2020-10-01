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

// GetTeamUserByFields retrieves teamUser by fields
func (database *Database) GetTeamUserByFields(connection *gorm.DB, teamUser *go_saas_model.TeamUser, fields map[string]interface{}) (*go_saas_model.TeamUser, error) {
	var query = connection

	for field, value := range fields {
		query = query.Where(fmt.Sprintf("team_users.%s = ?", field), value)
	}

	return teamUser, query.
		Preload("Team").
		Preload("User").
		Find(&teamUser).
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
		Preload("Team").
		Preload("User").
		Joins("JOIN users ON team_users.user_id = users.id").
		Where("team_users.team_id = ?", team.GetId()).
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

func (database *Database) UpdateRoleTeamUser(connection *gorm.DB, teamUser *go_saas_model.TeamUser) (*go_saas_model.TeamUser, error) {
	return teamUser, connection.
		Model(teamUser).
		Update(map[string]interface{}{
			"role": teamUser.GetRole(),
		}).
		Error
}

// DeleteTeamUser deletes teamUser
func (database *Database) DeleteTeamUser(connection *gorm.DB, teamUser *go_saas_model.TeamUser) error {
	return connection.
		Unscoped().
		Delete(teamUser).
		Error
}

// DeleteTeam deletes team and all their teamUsers
func (database *Database) DeleteTeam(connection *gorm.DB, team *go_saas_model.Team) error {
	return connection.
		Unscoped().
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

func (database *Database) IsModelTeam(connection *gorm.DB, team *go_saas_model.Team, model interface{}) (bool, error) {
	var count int

	return count == 1, connection.
		Model(model).
		Select("COUNT(*)").
		Where("team_id = ?", team.GetId()).
		Count(&count).
		Error
}
