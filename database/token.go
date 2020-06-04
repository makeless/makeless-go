package saas_database

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
	"sync"
)

func (database *Database) GetTokens(tokens []*go_saas_model.Token, userId *uint) ([]*go_saas_model.Token, error) {
	return tokens, database.GetConnection().
		Select([]string{
			"tokens.id",
			"tokens.note",
			"tokens.user_id",
			"tokens.team_id",
		}).
		Where("tokens.user_id = ? AND tokens.team_id IS NULL", userId).
		Order("tokens.id DESC").
		Find(&tokens).
		Error
}

func (database *Database) CreateToken(token *go_saas_model.Token, userId *uint) (*go_saas_model.Token, error) {
	token.UserId = userId
	token.TeamId = nil
	token.User = nil
	token.Team = nil
	token.RWMutex = new(sync.RWMutex)

	return token, database.GetConnection().
		Create(&token).
		Error
}

func (database *Database) DeleteToken(token *go_saas_model.Token, userId *uint) error {
	token.UserId = userId
	token.TeamId = nil
	token.RWMutex = new(sync.RWMutex)

	return database.GetConnection().
		Unscoped().
		Where("tokens.id = ? AND tokens.user_id = ? AND tokens.team_id IS NULL", token.GetId(), token.GetUserId()).
		Delete(&token).
		Error
}

func (database *Database) GetTokensTeam(tokens []*go_saas_model.Token, teamId uint, userId *uint) ([]*go_saas_model.Token, error) {
	return tokens, database.GetConnection().
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("users.id, users.name, users.username, users.email")
		}).
		Select([]string{
			"tokens.id",
			"tokens.note",
			"tokens.user_id",
			"tokens.team_id",
		}).
		Joins("JOIN teams ON teams.id = tokens.team_id").
		Where("tokens.team_id = ? AND teams.user_id = ?", teamId, userId).
		Order("tokens.id DESC").
		Find(&tokens).
		Error
}

// delete from multiple tables currently not supported by gorm
func (database *Database) DeleteTokenTeam(token *go_saas_model.Token, teamId *uint, userId *uint) error {
	token.UserId = userId
	token.TeamId = teamId
	token.RWMutex = new(sync.RWMutex)

	return database.GetConnection().
		Unscoped().
		Exec(
			"DELETE tokens FROM `tokens` JOIN teams ON teams.id = tokens.team_id WHERE tokens.id = ? AND tokens.team_id = ? AND teams.user_id = ?",
			token.GetId(),
			token.GetTeamId(),
			token.GetUserId(),
		).
		Error
}

func (database *Database) CreateTokenTeam(token *go_saas_model.Token, teamId *uint) (*go_saas_model.Token, error) {
	token.TeamId = teamId
	token.User = nil
	token.Team = nil

	return token, database.GetConnection().
		Create(&token).
		Error
}
