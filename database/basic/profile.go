package makeless_go_database_basic

import (
	"gorm.io/gorm"
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/struct"
)

func (database *Database) UpdateProfile(connection *gorm.DB, user *makeless_go_model.User, profile *_struct.Profile) (*makeless_go_model.User, error) {
	return user, connection.
		Model(&user).
		Update(map[string]interface{}{
			"name": profile.GetName(),
		}).
		Error
}
