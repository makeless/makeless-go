package makeless_go_database_basic

import (
	"github.com/makeless/makeless-go/model"
	"gorm.io/gorm"
)

func (database *Database) GetTokensTeam(connection *gorm.DB, team *makeless_go_model.Team, tokens []*makeless_go_model.Token) ([]*makeless_go_model.Token, error) {
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

// FIXME: Remove query
func (database *Database) CreateTokenTeam(connection *gorm.DB, token *makeless_go_model.Token) (*makeless_go_model.Token, error) {
	return token, connection.
		Create(token).
		Preload("User").
		First(token).
		Error
}

func (database *Database) DeleteTokenTeam(connection *gorm.DB, token *makeless_go_model.Token) error {
	return connection.
		Where("tokens.id = ? AND tokens.team_id = ?", token.GetId(), token.GetTeamId()).
		Delete(token).
		Error
}
