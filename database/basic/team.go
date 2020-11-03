package makeless_go_database_basic

import (
	"fmt"
	"github.com/makeless/makeless-go/model"
	"gorm.io/gorm"
)

// CreateTeam creates team
func (database *Database) CreateTeam(connection *gorm.DB, team *makeless_go_model.Team) (*makeless_go_model.Team, error) {
	return team, connection.
		Create(team).
		Error
}

// CreateTeam retrieves team
func (database *Database) GetTeam(connection *gorm.DB, team *makeless_go_model.Team) (*makeless_go_model.Team, error) {
	return team, connection.
		Preload("TeamUsers.Team").
		Preload("TeamUsers.User").
		Preload("TeamInvitations").
		First(team).
		Error
}

// AddTeamInvitations appends team invitations to a given team
func (database *Database) AddTeamInvitations(connection *gorm.DB, team *makeless_go_model.Team, teamInvitations []*makeless_go_model.TeamInvitation) (*makeless_go_model.Team, error) {
	return team, connection.
		Model(team).
		Association("TeamInvitations").
		Append(teamInvitations)
}

// GetTeamUserByFields retrieves teamUser by fields
func (database *Database) GetTeamUserByFields(connection *gorm.DB, teamUser *makeless_go_model.TeamUser, fields map[string]interface{}) (*makeless_go_model.TeamUser, error) {
	var query = connection

	for field, value := range fields {
		query = query.Where(fmt.Sprintf("team_users.%s = ?", field), value)
	}

	return teamUser, query.
		Preload("Team").
		Preload("User").
		First(&teamUser).
		Error
}

// GetTeamUsers retrieves team users by search
func (database *Database) GetTeamUsers(connection *gorm.DB, search string, teamUsers []*makeless_go_model.TeamUser, team *makeless_go_model.Team) ([]*makeless_go_model.TeamUser, error) {
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
func (database *Database) AddTeamUsers(connection *gorm.DB, teamUsers []*makeless_go_model.TeamUser, team *makeless_go_model.Team) error {
	return connection.
		Model(team).
		Association("TeamUsers").
		Append(teamUsers)
}

func (database *Database) UpdateRoleTeamUser(connection *gorm.DB, teamUser *makeless_go_model.TeamUser, role string) (*makeless_go_model.TeamUser, error) {
	return teamUser, connection.
		Model(teamUser).
		Updates(map[string]interface{}{
			"role": role,
		}).
		Error
}

// DeleteTeamUser deletes teamUser
func (database *Database) DeleteTeamUser(connection *gorm.DB, teamUser *makeless_go_model.TeamUser) error {
	return connection.
		Unscoped().
		Delete(teamUser).
		Error
}

// DeleteTeam deletes team and all their teamUsers and teamInvitations
func (database *Database) DeleteTeam(connection *gorm.DB, team *makeless_go_model.Team) error {
	return connection.
		Debug().
		Select("TeamUsers", "TeamInvitations").
		Delete(team).
		Error
}

// IsTeamUser checks if user is part of team
func (database *Database) IsTeamUser(connection *gorm.DB, team *makeless_go_model.Team, user *makeless_go_model.User) (bool, error) {
	var count int64

	return count == 1, connection.
		Raw("SELECT COUNT(*) FROM team_users WHERE team_users.team_id = ? AND team_users.user_id = ? LIMIT 1", team.GetId(), user.GetId()).
		Count(&count).
		Error
}

// IsTeamRole checks if user is part of team and has given role
func (database *Database) IsTeamRole(connection *gorm.DB, role string, team *makeless_go_model.Team, user *makeless_go_model.User) (bool, error) {
	var count int64

	return count == 1, connection.
		Raw("SELECT COUNT(*) FROM team_users WHERE team_users.team_id = ? AND team_users.user_id = ? AND team_users.role = ? LIMIT 1", team.GetId(), user.GetId(), role).
		Count(&count).
		Error
}

// IsTeamCreator checks if user is team creator
func (database *Database) IsTeamCreator(connection *gorm.DB, team *makeless_go_model.Team, user *makeless_go_model.User) (bool, error) {
	var count int64

	return count == 1, connection.
		Raw("SELECT COUNT(*) FROM teams WHERE teams.id = ? AND teams.user_id = ? LIMIT 1", team.GetId(), user.GetId()).
		Count(&count).
		Error
}

// IsNotTeamCreator checks if user is not team creator
func (database *Database) IsNotTeamCreator(connection *gorm.DB, team *makeless_go_model.Team, user *makeless_go_model.User) (bool, error) {
	var count int64

	return count == 0, connection.
		Raw("SELECT COUNT(*) FROM teams WHERE teams.id = ? AND teams.user_id = ? LIMIT 1", team.GetId(), user.GetId()).
		Count(&count).
		Error
}

func (database *Database) IsModelTeam(connection *gorm.DB, team *makeless_go_model.Team, model interface{}) (bool, error) {
	var count int64

	return count == 1, connection.
		Model(model).
		Select("COUNT(*)").
		Where("team_id = ?", team.GetId()).
		Count(&count).
		Error
}
