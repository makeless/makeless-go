package go_saas_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

type TeamInvitation interface {
	GetTeamInvitations(connection *gorm.DB, user *go_saas_model.User, teamInvitations []*go_saas_model.TeamInvitation) ([]*go_saas_model.TeamInvitation, error)
}
