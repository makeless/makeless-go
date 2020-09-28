package go_saas_database_basic

import (
	"fmt"
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
	"time"
)

func (database *Database) GetTeamInvitationByField(connection *gorm.DB, teamInvitation *go_saas_model.TeamInvitation, field string, value string) (*go_saas_model.TeamInvitation, error) {
	return teamInvitation, connection.
		Preload("Team").
		Preload("User").
		Where(
			fmt.Sprintf("team_invitations.%s = ? AND team_invitations.accepted = ? AND team_invitations.expire >= ?", field),
			value,
			false,
			time.Now(),
		).
		Find(teamInvitation).
		Error
}

func (database *Database) GetTeamInvitations(connection *gorm.DB, user *go_saas_model.User, teamInvitations []*go_saas_model.TeamInvitation) ([]*go_saas_model.TeamInvitation, error) {
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

func (database *Database) GetTeamInvitationsTeam(connection *gorm.DB, team *go_saas_model.Team, teamInvitations []*go_saas_model.TeamInvitation) ([]*go_saas_model.TeamInvitation, error) {
	return teamInvitations, connection.
		Preload("Team").
		Preload("User").
		Where("team_invitations.team_id = ? AND team_invitations.expire >= ? AND team_invitations.accepted = ?", team.GetId(), time.Now(), false).
		Order("team_invitations.id DESC").
		Find(&teamInvitations).
		Error
}

func (database *Database) AcceptTeamInvitation(connection *gorm.DB, teamInvitation *go_saas_model.TeamInvitation) (*go_saas_model.TeamInvitation, error) {
	return teamInvitation, connection.
		Model(teamInvitation).
		Update(map[string]interface{}{
			"accepted": true,
		}).
		Error
}

func (database *Database) DeleteTeamInvitation(connection *gorm.DB, teamInvitation *go_saas_model.TeamInvitation) (*go_saas_model.TeamInvitation, error) {
	return teamInvitation, connection.
		Delete(teamInvitation).
		Error
}
