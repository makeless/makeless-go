package go_saas_database

import "github.com/go-saas/go-saas/model"

type TokenTeam interface {
	GetTokensTeam(team *go_saas_model.Team, tokens []*go_saas_model.Token) ([]*go_saas_model.Token, error)
	DeleteTokenTeam(token *go_saas_model.Token) error
	CreateTokenTeam(token *go_saas_model.Token) (*go_saas_model.Token, error)
}
