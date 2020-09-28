package go_saas_database_basic

import (
	"fmt"
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

// GetUser retrieves user and all there team informations
func (database *Database) GetUser(connection *gorm.DB, user *go_saas_model.User) (*go_saas_model.User, error) {
	return user, connection.
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
func (database *Database) GetUserByField(connection *gorm.DB, user *go_saas_model.User, field string, value string) (*go_saas_model.User, error) {
	return user, connection.
		Where(fmt.Sprintf("%s = ?", field), value).
		Find(&user).
		Error
}

// CreateUser creates new user
func (database *Database) CreateUser(connection *gorm.DB, user *go_saas_model.User) (*go_saas_model.User, error) {
	return user, connection.
		Create(&user).
		Error
}

func (database *Database) UsersTeam(connection *gorm.DB, search string, users []*go_saas_model.User, team *go_saas_model.Team) ([]*go_saas_model.User, error) {
	var query = connection

	if search != "" {
		query = query.Where(
			"users.name LIKE ? OR users.email LIKE ?",
			fmt.Sprintf(`%s%s%s`, "%", search, "%"),
			fmt.Sprintf(`%s%s%s`, "%", search, "%"),
		)
	}

	return users, query.
		Select("users.id, users.name, users.email").
		Joins("JOIN team_users ON team_users.user_id = users.id AND team_users.team_id = ?", team.GetId()).
		Find(&users).
		Error
}

func (database *Database) IsModelUser(connection *gorm.DB, user *go_saas_model.User, model interface{}) (bool, error) {
	var count int

	return count == 1, connection.
		Model(model).
		Select("COUNT(*)").
		Where("user_id = ?", user.GetId()).
		Count(&count).
		Error
}
