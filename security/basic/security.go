package go_saas_security_basic

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/go-saas/go-saas/database"
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
	"sync"
)

type Security struct {
	Database go_saas_database.Database
	*sync.RWMutex
}

func (security *Security) GetDatabase() go_saas_database.Database {
	security.RLock()
	defer security.RUnlock()

	return security.Database
}

func (security *Security) Login(field string, value string, password string) (*go_saas_model.User, error) {
	var err error
	var user = &go_saas_model.User{
		RWMutex: new(sync.RWMutex),
	}

	if user, err = security.GetDatabase().GetUserByField(user, field, value); err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	if security.ComparePassword(*user.GetPassword(), password) != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return user, nil
}

func (security *Security) Register(user *go_saas_model.User) (*go_saas_model.User, error) {
	encrypted, err := security.EncryptPassword(*user.GetPassword())

	if err != nil {
		return nil, err
	}

	user.SetPassword(encrypted)
	if user, err = security.GetDatabase().CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (security *Security) UserExists(field string, value string) (bool, error) {
	_, err := security.GetDatabase().GetUserByField(new(go_saas_model.User), field, value)

	switch err {
	case gorm.ErrRecordNotFound:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}
