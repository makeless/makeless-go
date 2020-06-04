package saas_database

import (
	"fmt"
	"sync"

	"github.com/go-saas/go-saas/model"
)

func (database *Database) CreateTeam(team *go_saas_model.Team, userId *uint) (*go_saas_model.Team, error) {
	team.UserId = userId
	team.RWMutex = new(sync.RWMutex)

	return team, database.GetConnection().
		Model(&go_saas_model.User{
			Model: go_saas_model.Model{Id: *userId},
		}).
		Association("Teams").
		Append(team).
		Error
}

func (database *Database) LeaveTeam(team *go_saas_model.Team, userId *uint) error {
	team.UserId = userId
	team.RWMutex = new(sync.RWMutex)

	return database.GetConnection().
		Model(&go_saas_model.User{
			Model: go_saas_model.Model{Id: *userId},
		}).
		Association("Teams").
		Delete(team).
		Error
}

func (database *Database) DeleteTeam(team *go_saas_model.Team, userId *uint) error {
	team.UserId = userId
	team.RWMutex = new(sync.RWMutex)

	return database.GetConnection().
		Unscoped().
		Where("teams.id = ? AND teams.user_id = ?", team.GetId(), team.GetUserId()).
		Delete(team).
		Error
}

// GetUsersTeam queries all team users without any security restrictions (!)
func (database *Database) GetUsersTeam(search string, users []*go_saas_model.User, teamId *uint) ([]*go_saas_model.User, error) {
	return users, database.GetConnection().
		Select("users.id, users.name, users.username, users.email").
		Where(
			"users.name LIKE ? OR users.email LIKE ?",
			fmt.Sprintf(`%s%s%s`, "%", search, "%"),
			fmt.Sprintf(`%s%s%s`, "%", search, "%"),
		).
		Model(&go_saas_model.Team{
			Model: go_saas_model.Model{Id: *teamId},
		}).
		Related(&users, "Users").
		Error
}
