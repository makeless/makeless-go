package makeless_go_database_basic

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/makeless/makeless-go/model"
)

// GetUser retrieves user and all there team informations
func (database *Database) GetUser(connection *gorm.DB, user *makeless_go_model.User) (*makeless_go_model.User, error) {
	return user, connection.
		Preload("EmailVerification").
		Preload("TeamUsers", func(db *gorm.DB) *gorm.DB {
			return db.Where("team_users.user_id = ?", user.GetId())
		}).
		Preload("TeamUsers.Team").
		Preload("TeamUsers.User").
		First(&user).
		Error
}

// GetUserByField retrieves user by specific field
// Mostly used for login mechanisms
// Do not use this for outputs
func (database *Database) GetUserByField(connection *gorm.DB, user *makeless_go_model.User, field string, value string) (*makeless_go_model.User, error) {
	return user, connection.
		Preload("EmailVerification").
		Where(fmt.Sprintf("%s = ?", field), value).
		Find(&user).
		Error
}

// CreateUser creates new user
func (database *Database) CreateUser(connection *gorm.DB, user *makeless_go_model.User) (*makeless_go_model.User, error) {
	return user, connection.
		Create(&user).
		Error
}

func (database *Database) IsModelUser(connection *gorm.DB, user *makeless_go_model.User, model interface{}) (bool, error) {
	var count int

	return count == 1, connection.
		Model(model).
		Select("COUNT(*)").
		Where("user_id = ?", user.GetId()).
		Count(&count).
		Error
}
