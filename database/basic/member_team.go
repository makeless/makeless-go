package go_saas_database_basic

import (
	"fmt"
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

func (database *Database) MembersTeam(connection *gorm.DB, search string, users []*go_saas_model.User, team *go_saas_model.Team) ([]*go_saas_model.User, error) {
	var query = connection

	if search != "" {
		query = query.Where(
			"users.name LIKE ? OR users.email LIKE ?",
			fmt.Sprintf(`%s%s%s`, "%", search, "%"),
			fmt.Sprintf(`%s%s%s`, "%", search, "%"),
		)
	}

	return users, query.
		Select("users.id, users.name, users.email").
		Joins("JOIN team_users ON team_users.user_id = users.id AND team_users.team_id = ?", team.GetId()).
		Find(&users).
		Error
}
