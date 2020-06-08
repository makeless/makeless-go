package go_saas_basic_database

import (
	"fmt"
	"github.com/go-saas/go-saas/model"
)

func (database *Database) GetMembersTeam(team *go_saas_model.Team, users []*go_saas_model.User) ([]*go_saas_model.User, error) {
	return users, database.GetConnection().
		Model(team).
		Related(&users, "Users").
		Error
}

func (database *Database) RemoveMemberTeam(user *go_saas_model.User, team *go_saas_model.Team) error {
	return database.GetConnection().
		Model(user).
		Association("Teams").
		Delete(team).
		Error
}

func (database *Database) SearchMembersTeam(search string, team *go_saas_model.Team, users []*go_saas_model.User) ([]*go_saas_model.User, error) {
	return users, database.GetConnection().
		Select("users.id, users.name, users.email").
		Where(
			"users.name LIKE ? OR users.email LIKE ?",
			fmt.Sprintf(`%s%s%s`, "%", search, "%"),
			fmt.Sprintf(`%s%s%s`, "%", search, "%"),
		).
		Model(team).
		Related(&users, "Users").
		Error
}
