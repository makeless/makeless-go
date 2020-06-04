package go_saas_database

import "github.com/go-saas/go-saas/model"

type Team interface {
	CreateTeam(user *go_saas_model.User, team *go_saas_model.Team) (*go_saas_model.Team, error)
	LeaveTeam(user *go_saas_model.User, team *go_saas_model.Team) error
	DeleteTeam(team *go_saas_model.Team) error
}
