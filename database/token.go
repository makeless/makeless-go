package makeless_go_database

import (
	"github.com/makeless/makeless-go/model"
	"github.com/jinzhu/gorm"
)

type Token interface {
	GetTokens(connection *gorm.DB, user *makeless_go_model.User, tokens []*makeless_go_model.Token) ([]*makeless_go_model.Token, error)
	CreateToken(connection *gorm.DB, token *makeless_go_model.Token) (*makeless_go_model.Token, error)
	DeleteToken(connection *gorm.DB, token *makeless_go_model.Token) error

	GetTokensTeam(connection *gorm.DB, team *makeless_go_model.Team, tokens []*makeless_go_model.Token) ([]*makeless_go_model.Token, error)
	DeleteTokenTeam(connection *gorm.DB, token *makeless_go_model.Token) error
	CreateTokenTeam(connection *gorm.DB, token *makeless_go_model.Token) (*makeless_go_model.Token, error)
}
