package saas_database

import (
	"github.com/loeffel-io/go-saas/model"
	"sync"
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

func (database *Database) DeleteTeam(team *saas_model.Team, userId *uint) error {
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