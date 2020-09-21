package go_saas_database_basic

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
	"time"
)

func (database *Database) IsTeamInvitation(connection *gorm.DB, teamInvitation *go_saas_model.TeamInvitation) (bool, error) {
	var count int

	return count == 1, connection.
		Model(teamInvitation).
		Select("COUNT(*)").
		Where(
			"team_invitations.team_id = ? AND team_invitations.email = ? AND team_invitations.token = ? AND team_invitations.accepted = ? AND team_invitations.expire >= ?",
			teamInvitation.GetTeamId(),
			teamInvitation.GetEmail(),
			teamInvitation.GetToken(),
			false,
			time.Now(),
		).
		Count(&count).
		Error
}

func (database *Database) GetTeamInvitations(connection *gorm.DB, user *go_saas_model.User, teamInvitations []*go_saas_model.TeamInvitation) ([]*go_saas_model.TeamInvitation, error) {
	return teamInvitations, connection.
		Preload("Team").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("users.id, users.name, users.email")
		}).
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
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("users.id, users.name, users.email")
		}).
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

func (database *Database) DeclineTeamInvitation(connection *gorm.DB, teamInvitation *go_saas_model.TeamInvitation) (*go_saas_model.TeamInvitation, error) {
	return teamInvitation, connection.
		Delete(teamInvitation).
		Error
}
