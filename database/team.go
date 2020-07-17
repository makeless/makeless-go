package go_saas_database

import "github.com/go-saas/go-saas/model"

type Team interface {
	CreateTeam(team *go_saas_model.Team, teamUser *go_saas_model.TeamUser) (*go_saas_model.Team, error)
	DeleteTeamUsers(user *go_saas_model.User, team *go_saas_model.Team) error
	DeleteTeam(team *go_saas_model.Team) error
	IsTeamMember(team *go_saas_model.Team, user *go_saas_model.User) (bool, error)
	IsTeamOwner(team *go_saas_model.Team, user *go_saas_model.User) (bool, error)
}
