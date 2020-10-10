package makeless_go_security_basic

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/jinzhu/gorm"
	"github.com/makeless/makeless-go/database"
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/security"
	"sync"
)

type Security struct {
	Database makeless_go_database.Database
	*sync.RWMutex
}

func (security *Security) GetDatabase() makeless_go_database.Database {
	security.RLock()
	defer security.RUnlock()

	return security.Database
}

func (security *Security) Login(connection *gorm.DB, field string, value string, password string) (*makeless_go_model.User, error) {
	var err error
	var user = &makeless_go_model.User{
		RWMutex: new(sync.RWMutex),
	}

	if user, err = security.GetDatabase().GetUserByField(connection, user, field, value); err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	if security.ComparePassword(*user.GetPassword(), password) != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	if user.GetEmailVerification() != nil {
		user.GetEmailVerification().RWMutex = new(sync.RWMutex)
	}

	return user, nil
}

func (security *Security) Register(connection *gorm.DB, user *makeless_go_model.User) (*makeless_go_model.User, error) {
	exists, err := security.UserExists(connection, "email", *user.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, makeless_go_security.UserAlreadyExist
	}

	encrypted, err := security.EncryptPassword(*user.GetPassword())
	if err != nil {
		return nil, err
	}

	user.SetPassword(encrypted)
	if user, err = security.GetDatabase().CreateUser(connection, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (security *Security) UserExists(connection *gorm.DB, field string, value string) (bool, error) {
	_, err := security.GetDatabase().GetUserByField(connection, new(makeless_go_model.User), field, value)

	switch err {
	case gorm.ErrRecordNotFound:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}

func (security *Security) IsModelUser(connection *gorm.DB, userId uint, model interface{}) (bool, error) {
	var user = &makeless_go_model.User{
		Model:   makeless_go_model.Model{Id: userId},
		RWMutex: new(sync.RWMutex),
	}

	return security.GetDatabase().IsModelUser(connection, user, model)
}

func (security *Security) IsModelTeam(connection *gorm.DB, teamId uint, model interface{}) (bool, error) {
	var team = &makeless_go_model.Team{
		Model:   makeless_go_model.Model{Id: teamId},
		RWMutex: new(sync.RWMutex),
	}

	return security.GetDatabase().IsModelTeam(connection, team, model)
}
