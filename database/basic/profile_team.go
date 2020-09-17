package go_saas_database_basic

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

func (database *Database) UpdateProfileTeam(connection *gorm.DB, team *go_saas_model.Team) (*go_saas_model.Team, error) {
	return team, connection.
		Model(&team).
		Update(map[string]interface{}{
			"name": team.Name,
		}).
		Error
}
