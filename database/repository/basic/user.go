package makeless_go_repository_basic

import (
	"fmt"
	"github.com/makeless/makeless-go/database/model"
	"gorm.io/gorm"
)

type UserRepository struct {
}

// GetUser retrieves user and all there team informations
func (userRepository *UserRepository) GetUser(connection *gorm.DB, user *makeless_go_model.User) (*makeless_go_model.User, error) {
	return user, connection.
		Preload("EmailVerification").
		Preload("TeamUsers", func(db *gorm.DB) *gorm.DB {
			return db.Where("team_users.user_id = ?", user.Id)
		}).
		Preload("TeamUsers.Team").
		Preload("TeamUsers.User").
		First(&user).
		Error
}

// GetUserByField retrieves user by specific field
// Mostly used for login mechanisms
func (userRepository *UserRepository) GetUserByField(connection *gorm.DB, user *makeless_go_model.User, field string, value string) (*makeless_go_model.User, error) {
	return user, connection.
		Preload("EmailVerification").
		Where(fmt.Sprintf("%s = ?", field), value).
		First(&user).
		Error
}

// CreateUser creates new user
func (userRepository *UserRepository) CreateUser(connection *gorm.DB, user *makeless_go_model.User) (*makeless_go_model.User, error) {
	return user, connection.
		Create(&user).
		Error
}

// IsModelUser checks if model belongs to user
func (userRepository *UserRepository) IsModelUser(connection *gorm.DB, user *makeless_go_model.User, model interface{}) (bool, error) {
	var count int64

	return count >= 1, connection.
		Model(model).
		Select("COUNT(*)").
		Where("user_id = ?", user.Id).
		Count(&count).
		Error
}
