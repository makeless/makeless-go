package go_saas_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

type Team interface {
	CreateTeam(connection *gorm.DB, team *go_saas_model.Team) (*go_saas_model.Team, error)
	DeleteTeam(connection *gorm.DB, team *go_saas_model.Team) error
	GetTeam(connection *gorm.DB, team *go_saas_model.Team) (*go_saas_model.Team, error)
	AddTeamInvitations(connection *gorm.DB, team *go_saas_model.Team, teamInvitations []*go_saas_model.TeamInvitation) (*go_saas_model.Team, error)
	GetTeamUserByFields(connection *gorm.DB, teamUser *go_saas_model.TeamUser, fields map[string]interface{}) (*go_saas_model.TeamUser, error)
	GetTeamUsers(connection *gorm.DB, search string, teamUsers []*go_saas_model.TeamUser, team *go_saas_model.Team) ([]*go_saas_model.TeamUser, error)
	AddTeamUsers(connection *gorm.DB, teamUsers []*go_saas_model.TeamUser, team *go_saas_model.Team) error
	UpdateRoleTeamUser(connection *gorm.DB, teamUser *go_saas_model.TeamUser) (*go_saas_model.TeamUser, error)
	DeleteTeamUser(connection *gorm.DB, teamUser *go_saas_model.TeamUser) error
	IsTeamUser(connection *gorm.DB, team *go_saas_model.Team, user *go_saas_model.User) (bool, error)
	IsTeamRole(connection *gorm.DB, role string, team *go_saas_model.Team, user *go_saas_model.User) (bool, error)
	IsTeamCreator(connection *gorm.DB, team *go_saas_model.Team, user *go_saas_model.User) (bool, error)
	IsNotTeamCreator(connection *gorm.DB, team *go_saas_model.Team, user *go_saas_model.User) (bool, error)
	IsModelTeam(connection *gorm.DB, team *go_saas_model.Team, model interface{}) (bool, error)
}
