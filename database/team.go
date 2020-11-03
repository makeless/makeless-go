package makeless_go_database

import (
	"github.com/makeless/makeless-go/model"
	"gorm.io/gorm"
)

type Team interface {
	CreateTeam(connection *gorm.DB, team *makeless_go_model.Team) (*makeless_go_model.Team, error)
	DeleteTeam(connection *gorm.DB, team *makeless_go_model.Team) error
	GetTeam(connection *gorm.DB, team *makeless_go_model.Team) (*makeless_go_model.Team, error)
	AddTeamInvitations(connection *gorm.DB, team *makeless_go_model.Team, teamInvitations []*makeless_go_model.TeamInvitation) (*makeless_go_model.Team, error)
	GetTeamUserByFields(connection *gorm.DB, teamUser *makeless_go_model.TeamUser, fields map[string]interface{}) (*makeless_go_model.TeamUser, error)
	GetTeamUsers(connection *gorm.DB, search string, teamUsers []*makeless_go_model.TeamUser, team *makeless_go_model.Team) ([]*makeless_go_model.TeamUser, error)
	AddTeamUsers(connection *gorm.DB, teamUsers []*makeless_go_model.TeamUser, team *makeless_go_model.Team) error
	UpdateRoleTeamUser(connection *gorm.DB, teamUser *makeless_go_model.TeamUser, role string) (*makeless_go_model.TeamUser, error)
	DeleteTeamUser(connection *gorm.DB, teamUser *makeless_go_model.TeamUser) error
	IsTeamUser(connection *gorm.DB, team *makeless_go_model.Team, user *makeless_go_model.User) (bool, error)
	IsTeamRole(connection *gorm.DB, role string, team *makeless_go_model.Team, user *makeless_go_model.User) (bool, error)
	IsTeamCreator(connection *gorm.DB, team *makeless_go_model.Team, user *makeless_go_model.User) (bool, error)
	IsNotTeamCreator(connection *gorm.DB, team *makeless_go_model.Team, user *makeless_go_model.User) (bool, error)
	IsModelTeam(connection *gorm.DB, team *makeless_go_model.Team, model interface{}) (bool, error)
}
