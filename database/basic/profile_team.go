package makeless_go_database_basic

import (
	"gorm.io/gorm"
	"github.com/makeless/makeless-go/model"
)

func (database *Database) UpdateProfileTeam(connection *gorm.DB, team *makeless_go_model.Team) (*makeless_go_model.Team, error) {
	return team, connection.
		Model(&team).
		Update(map[string]interface{}{
			"name": team.Name,
		}).
		Error
}
