package saas_database

import (
	"github.com/loeffel-io/go-saas/model"
)

func (database *Database) GetTokens(userId uint) ([]*saas_model.Token, error) {
	var tokens []*saas_model.Token

	return tokens, database.GetConnection().
		Where("tokens.user_id = ?", userId).
		Find(&tokens).
		Error
}

func (database *Database) CreateToken(token *saas_model.Token) (*saas_model.Token, error) {
	return token, database.GetConnection().Create(&token).Error
}

func (database *Database) DeleteToken(token *saas_model.Token) error {
	return database.GetConnection().
		Unscoped().
		Where("tokens.id = ? AND tokens.token = ? AND tokens.user_id = ?", token.GetId(), token.GetToken(), token.GetUserId()).
		Delete(&token).
		Error
}
