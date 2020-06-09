package go_saas_database

import "github.com/go-saas/go-saas/model"

type Token interface {
	GetToken(token *go_saas_model.Token, value string) (*go_saas_model.Token, error)
	GetTokens(user *go_saas_model.User, tokens []*go_saas_model.Token) ([]*go_saas_model.Token, error)
	CreateToken(token *go_saas_model.Token) (*go_saas_model.Token, error)
	DeleteToken(token *go_saas_model.Token) error

	GetTokensTeam(team *go_saas_model.Team, tokens []*go_saas_model.Token) ([]*go_saas_model.Token, error)
	DeleteTokenTeam(token *go_saas_model.Token) error
	CreateTokenTeam(token *go_saas_model.Token) (*go_saas_model.Token, error)
}
