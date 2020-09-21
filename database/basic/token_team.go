package go_saas_database_basic

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

func (database *Database) GetTokensTeam(connection *gorm.DB, team *go_saas_model.Team, tokens []*go_saas_model.Token) ([]*go_saas_model.Token, error) {
	return tokens, connection.
		Preload("User").
		Select([]string{
			"tokens.id",
			"tokens.note",
			"tokens.user_id",
			"tokens.team_id",
			"CONCAT(REPEAT('X', CHAR_LENGTH(tokens.token) - 4),SUBSTRING(tokens.token, -4)) as token",
		}).
		Joins("JOIN teams ON teams.id = tokens.team_id").
		Where("tokens.team_id = ?", team.GetId()).
		Order("tokens.id DESC").
		Find(&tokens).
		Error
}

func (database *Database) CreateTokenTeam(connection *gorm.DB, token *go_saas_model.Token) (*go_saas_model.Token, error) {
	return token, connection.
		Create(token).
		Preload("User").
		Find(token).
		Error
}

func (database *Database) DeleteTokenTeam(connection *gorm.DB, token *go_saas_model.Token) error {
	return connection.
		Unscoped().
		Where("tokens.id = ? AND tokens.team_id = ?", token.GetId(), token.GetTeamId()).
		Delete(&token).
		Error
}
