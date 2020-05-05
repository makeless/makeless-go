package saas_database

import (
	"github.com/loeffel-io/go-saas/model"
	"sync"
)

func (database *Database) GetTokens(tokens []*saas_model.Token, userId *uint) ([]*saas_model.Token, error) {
	return tokens, database.GetConnection().
		Select([]string{
			"tokens.id",
			"CONCAT(REPEAT('x', CHAR_LENGTH(tokens.token)-4), RIGHT(tokens.token, 4)) as token",
			"tokens.note",
			"tokens.user_id",
			"tokens.team_id",
		}).
		Where("tokens.user_id = ? AND tokens.team_id IS NULL", userId).
		Order("tokens.id DESC").
		Find(&tokens).
		Error
}

func (database *Database) CreateToken(token *saas_model.Token, userId *uint) (*saas_model.Token, error) {
	token.UserId = userId
	token.TeamId = nil
	token.RWMutex = new(sync.RWMutex)

	return token, database.GetConnection().
		Create(&token).
		Error
}

func (database *Database) DeleteToken(token *saas_model.Token, userId *uint) error {
	token.UserId = userId
	token.TeamId = nil
	token.RWMutex = new(sync.RWMutex)

	return database.GetConnection().
		Unscoped().
		Where("tokens.id = ? AND tokens.user_id = ? AND tokens.team_id IS NULL", token.GetId(), token.GetUserId()).
		Delete(&token).
		Error
}
