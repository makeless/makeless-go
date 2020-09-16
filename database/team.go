package go_saas_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

type Team interface {
	CreateTeam(connection *gorm.DB, team *go_saas_model.Team) (*go_saas_model.Team, error)
	AddTeamInvitations(connection *gorm.DB, team *go_saas_model.Team, teamInvitations []*go_saas_model.TeamInvitation) (*go_saas_model.Team, error)
	DeleteTeamUser(connection *gorm.DB, user *go_saas_model.User, team *go_saas_model.Team) error
	DeleteTeam(connection *gorm.DB, user *go_saas_model.User, team *go_saas_model.Team) error
	IsTeamUser(connection *gorm.DB, team *go_saas_model.Team, user *go_saas_model.User) (bool, error)
	IsTeamRole(connection *gorm.DB, role string, team *go_saas_model.Team, user *go_saas_model.User) (bool, error)
	IsTeamCreator(connection *gorm.DB, team *go_saas_model.Team, user *go_saas_model.User) (bool, error)
}
