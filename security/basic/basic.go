package go_saas_basic_security

import (
	"fmt"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/go-saas/go-saas/database/go_saas_basic_database"
	"github.com/go-saas/go-saas/model"
	"sync"
)

type Security struct {
	Database *go_saas_basic_database.Database
	*sync.RWMutex
}

func (security *Security) Login(field string, id string, password string) (*go_saas_model.User, error) {
	var user = &go_saas_model.User{
		RWMutex: new(sync.RWMutex),
	}

	if err := security.getDatabase().GetConnection().Where(fmt.Sprintf("%s = ?", field), id).Find(&user).Error; err != nil {
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

	user.SetPassword(string(encrypted))

	if err := security.getDatabase().GetConnection().Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
