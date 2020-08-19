package go_saas_database_basic

import (
	"fmt"
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

// GetUser retrieves user and all there team informations
func (database *Database) GetUser(user *go_saas_model.User) (*go_saas_model.User, error) {
	return user, database.GetConnection().
		Select("users.id, users.name, users.email").
		Preload("TeamUsers", func(db *gorm.DB) *gorm.DB {
			return db.Where("team_users.user_id = ?", user.GetId())
		}).
		Preload("TeamUsers.Team").
		Preload("TeamUsers.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("users.id, users.name, users.email")
		}).
		First(&user).
		Error
}

// GetUserByField retrieves user by specific field
// Mostly used for login mechanisms
// Do not use this for outputs
func (database *Database) GetUserByField(user *go_saas_model.User, field string, value string) (*go_saas_model.User, error) {
	return user, database.GetConnection().
		Where(fmt.Sprintf("%s = ?", field), value).
		Find(&user).
		Error
}

// CreateUser creates new user
func (database *Database) CreateUser(user *go_saas_model.User) (*go_saas_model.User, error) {
	return user, database.GetConnection().
		Create(&user).
		Error
}
