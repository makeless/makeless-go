package saas_database

import (
	"github.com/loeffel-io/go-saas/model"
	"sync"
)

func (database *Database) GetUser(userId uint) (*saas_model.User, error) {
	var user = &saas_model.User{
		RWMutex: new(sync.RWMutex),
	}

	return user, database.GetConnection().
		Select("users.id, users.name, users.username, users.email").
		Preload("Teams").
		Where("users.id = ?", userId).
		First(&user).
		Error
}
