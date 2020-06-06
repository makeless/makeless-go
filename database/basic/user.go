package go_saas_basic_database

import (
	"fmt"
	"github.com/go-saas/go-saas/model"
)

func (database *Database) GetUser(user *go_saas_model.User) (*go_saas_model.User, error) {
	return user, database.GetConnection().
		Select("users.id, users.name, users.email").
		Preload("Teams").
		First(&user).
		Error
}

func (database *Database) GetUserByField(user *go_saas_model.User, field string, value string) (*go_saas_model.User, error) {
	return user, database.GetConnection().
		Where(fmt.Sprintf("%s = ?", field), value).
		Find(&user).
		Error
}

func (database *Database) CreateUser(user *go_saas_model.User) (*go_saas_model.User, error) {
	return user, database.GetConnection().
		Create(&user).
		Error
}