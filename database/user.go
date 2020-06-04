package saas_database

import (
	"github.com/go-saas/go-saas/model"
	"sync"
)

func (database *Database) GetUser(userId uint) (*go_saas_model.User, error) {
	var user = &go_saas_model.User{
		RWMutex: new(sync.RWMutex),
	}

	return user, database.GetConnection().
		Select("users.id, users.name, users.username, users.email").
		Preload("Teams").
		Where("users.id = ?", userId).
		First(&user).
		Error
}
