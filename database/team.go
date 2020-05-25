package saas_database

import (
	"sync"

	saas_model "github.com/go-saas/go-saas/model"
)

func (database *Database) CreateTeam(team *saas_model.Team, userId *uint) (*saas_model.Team, error) {
	team.UserId = userId
	team.RWMutex = new(sync.RWMutex)

	return team, database.GetConnection().
		Model(&saas_model.User{
			Model: saas_model.Model{Id: *userId},
		}).
		Association("Teams").
		Append(team).
		Error
}

func (database *Database) LeaveTeam(team *saas_model.Team, userId *uint) error {
	team.UserId = userId
	team.RWMutex = new(sync.RWMutex)

	return database.GetConnection().
		Model(&saas_model.User{
			Model: saas_model.Model{Id: *userId},
		}).
		Association("Teams").
		Delete(team).
		Error
}

func (database *Database) DeleteTeam(team *saas_model.Team, userId *uint) error {
	team.UserId = userId
	team.RWMutex = new(sync.RWMutex)

	return database.GetConnection().
		Unscoped().
		Where("teams.id = ? AND teams.user_id = ?", team.GetId(), team.GetUserId()).
		Delete(team).
		Error
}
