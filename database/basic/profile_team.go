package makeless_go_database_basic

import (
	"github.com/makeless/makeless-go/model"
	"gorm.io/gorm"
)

func (database *Database) UpdateProfileTeam(connection *gorm.DB, team *makeless_go_model.Team) (*makeless_go_model.Team, error) {
	return team, connection.
		Model(&team).
		Updates(map[string]interface{}{
			"name": team.Name,
		}).
		Error
}
