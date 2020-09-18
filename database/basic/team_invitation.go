package go_saas_database_basic

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
	"time"
)

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
