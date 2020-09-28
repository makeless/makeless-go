package go_saas_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

type TeamInvitation interface {
	GetTeamInvitationByField(connection *gorm.DB, teamInvitation *go_saas_model.TeamInvitation, field string, value string) (*go_saas_model.TeamInvitation, error)
	GetTeamInvitations(connection *gorm.DB, user *go_saas_model.User, teamInvitations []*go_saas_model.TeamInvitation) ([]*go_saas_model.TeamInvitation, error)
	GetTeamInvitationsTeam(connection *gorm.DB, team *go_saas_model.Team, teamInvitations []*go_saas_model.TeamInvitation) ([]*go_saas_model.TeamInvitation, error)
	AcceptTeamInvitation(connection *gorm.DB, teamInvitation *go_saas_model.TeamInvitation) (*go_saas_model.TeamInvitation, error)
	DeleteTeamInvitation(connection *gorm.DB, teamInvitation *go_saas_model.TeamInvitation) (*go_saas_model.TeamInvitation, error)
}
