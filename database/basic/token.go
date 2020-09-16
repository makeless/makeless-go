package go_saas_database_basic

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

func (database *Database) GetTokens(connection *gorm.DB, user *go_saas_model.User, tokens []*go_saas_model.Token) ([]*go_saas_model.Token, error) {
	return tokens, connection.
		Select([]string{
			"tokens.id",
			"tokens.note",
			"tokens.user_id",
			"tokens.team_id",
			"CONCAT(REPEAT('X', CHAR_LENGTH(tokens.token) - 4),SUBSTRING(tokens.token, -4)) as token",
		}).
		Where("tokens.user_id = ? AND tokens.team_id IS NULL", user.GetId()).
		Order("tokens.id DESC").
		Find(&tokens).
		Error
}

func (database *Database) CreateToken(connection *gorm.DB, token *go_saas_model.Token) (*go_saas_model.Token, error) {
	return token, connection.
		Create(&token).
		Error
}

func (database *Database) DeleteToken(connection *gorm.DB, token *go_saas_model.Token) error {
	return connection.
		Unscoped().
		Where("tokens.id = ? AND tokens.user_id = ? AND tokens.team_id IS NULL", token.GetId(), token.GetUserId()).
		Delete(&token).
		Error
}
