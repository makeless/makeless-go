package go_saas_basic_database

import (
	"github.com/go-saas/go-saas/model"
)

func (database *Database) GetUser(user *go_saas_model.User) (*go_saas_model.User, error) {
	return user, database.GetConnection().
		Select("users.id, users.name, users.username, users.email").
		Preload("Teams").
		First(&user).
		Error
}
