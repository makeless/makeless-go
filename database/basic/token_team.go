package go_saas_basic_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

func (database *Database) GetTokensTeam(team *go_saas_model.Team, tokens []*go_saas_model.Token) ([]*go_saas_model.Token, error) {
	return tokens, database.GetConnection().
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("users.id, users.name, users.email")
		}).
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

func (database *Database) CreateTokenTeam(token *go_saas_model.Token) (*go_saas_model.Token, error) {
	return token, database.GetConnection().
		Create(token).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("users.id, users.name, users.email")
		}).
		Find(token).
		Error
}

func (database *Database) DeleteTokenTeam(token *go_saas_model.Token) error {
	return database.GetConnection().
		Unscoped().
		Where("tokens.id = ? AND tokens.team_id = ?", token.GetId(), token.GetTeamId()).
		Delete(&token).
		Error
}
