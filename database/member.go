package go_saas_database

import "github.com/go-saas/go-saas/model"

type Member interface {
	MembersTeam(search string, users []*go_saas_model.User, team *go_saas_model.Team) ([]*go_saas_model.User, error)
}
