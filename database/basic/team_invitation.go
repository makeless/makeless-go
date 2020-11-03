package makeless_go_database_basic

import (
	"fmt"
	"github.com/makeless/makeless-go/model"
	"gorm.io/gorm"
	"time"
)

func (database *Database) GetTeamInvitationByField(connection *gorm.DB, teamInvitation *makeless_go_model.TeamInvitation, field string, value string) (*makeless_go_model.TeamInvitation, error) {
	return teamInvitation, connection.
		Preload("Team").
		Preload("User").
		Where(
			fmt.Sprintf("team_invitations.%s = ? AND team_invitations.accepted = ? AND team_invitations.expire >= ?", field),
			value,
			false,
			time.Now(),
		).
		First(teamInvitation).
		Error
}

func (database *Database) GetTeamInvitations(connection *gorm.DB, user *makeless_go_model.User, teamInvitations []*makeless_go_model.TeamInvitation) ([]*makeless_go_model.TeamInvitation, error) {
	return teamInvitations, connection.
		Preload("Team").
		Preload("User").
		Joins("JOIN users ON users.email = team_invitations.email").
		Where("users.id = ?", user.GetId()).
		Where("team_invitations.expire >= ? AND team_invitations.accepted = ?", time.Now(), false).
		Order("team_invitations.id DESC").
		Find(&teamInvitations).
		Error
}

func (database *Database) GetTeamInvitationsTeam(connection *gorm.DB, team *makeless_go_model.Team, teamInvitations []*makeless_go_model.TeamInvitation) ([]*makeless_go_model.TeamInvitation, error) {
	return teamInvitations, connection.
		Preload("Team").
		Preload("User").
		Where("team_invitations.team_id = ? AND team_invitations.expire >= ? AND team_invitations.accepted = ?", team.GetId(), time.Now(), false).
		Order("team_invitations.id DESC").
		Find(&teamInvitations).
		Error
}

func (database *Database) AcceptTeamInvitation(connection *gorm.DB, teamInvitation *makeless_go_model.TeamInvitation) (*makeless_go_model.TeamInvitation, error) {
	return teamInvitation, connection.
		Model(teamInvitation).
		Updates(map[string]interface{}{
			"accepted": true,
		}).
		Error
}

func (database *Database) DeleteTeamInvitation(connection *gorm.DB, teamInvitation *makeless_go_model.TeamInvitation) (*makeless_go_model.TeamInvitation, error) {
	return teamInvitation, connection.
		Unscoped().
		Delete(teamInvitation).
		Error
}
