package go_saas_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

type Token interface {
	GetTokens(connection *gorm.DB, user *go_saas_model.User, tokens []*go_saas_model.Token) ([]*go_saas_model.Token, error)
	CreateToken(connection *gorm.DB, token *go_saas_model.Token) (*go_saas_model.Token, error)
	DeleteToken(connection *gorm.DB, token *go_saas_model.Token) error

	GetTokensTeam(connection *gorm.DB, team *go_saas_model.Team, tokens []*go_saas_model.Token) ([]*go_saas_model.Token, error)
	DeleteTokenTeam(connection *gorm.DB, token *go_saas_model.Token) error
	CreateTokenTeam(connection *gorm.DB, token *go_saas_model.Token) (*go_saas_model.Token, error)
}
