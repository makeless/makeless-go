package go_saas_basic_database

import (
	"fmt"
	"github.com/go-saas/go-saas/model"
)

func (database *Database) MembersTeam(search string, users []*go_saas_model.User, team *go_saas_model.Team) ([]*go_saas_model.User, error) {
	var query = database.GetConnection()

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

func (database *Database) RemoveMemberTeam(user *go_saas_model.User, team *go_saas_model.Team) error {
	return database.GetConnection().
		Model(user).
		Association("Teams").
		Delete(team).
		Error
}
