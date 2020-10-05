package makeless_go_database

import (
	"github.com/makeless/makeless-go/model"
	"github.com/jinzhu/gorm"
)

type TeamInvitation interface {
	GetTeamInvitationByField(connection *gorm.DB, teamInvitation *makeless_go_model.TeamInvitation, field string, value string) (*makeless_go_model.TeamInvitation, error)
	GetTeamInvitations(connection *gorm.DB, user *makeless_go_model.User, teamInvitations []*makeless_go_model.TeamInvitation) ([]*makeless_go_model.TeamInvitation, error)
	GetTeamInvitationsTeam(connection *gorm.DB, team *makeless_go_model.Team, teamInvitations []*makeless_go_model.TeamInvitation) ([]*makeless_go_model.TeamInvitation, error)
	AcceptTeamInvitation(connection *gorm.DB, teamInvitation *makeless_go_model.TeamInvitation) (*makeless_go_model.TeamInvitation, error)
	DeleteTeamInvitation(connection *gorm.DB, teamInvitation *makeless_go_model.TeamInvitation) (*makeless_go_model.TeamInvitation, error)
}
