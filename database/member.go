package go_saas_database

import "github.com/go-saas/go-saas/model"

type Member interface {
	GetMembersTeam(team *go_saas_model.Team, users []*go_saas_model.User) ([]*go_saas_model.User, error)
	RemoveMemberTeam(user *go_saas_model.User, team *go_saas_model.Team) error
	SearchMembersTeam(search string, team *go_saas_model.Team, users []*go_saas_model.User) ([]*go_saas_model.User, error)
}
