package saas_security_basic

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/loeffel-io/go-saas/database"
	"github.com/loeffel-io/go-saas/model"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type Basic struct {
	Database *saas_database.Database
	*sync.RWMutex
}

func (basic *Basic) getDatabase() *saas_database.Database {
	basic.RLock()
	defer basic.RUnlock()

	return basic.Database
}

func (basic *Basic) Login(username string, password string) (*saas_model.User, error) {
	var user = &saas_model.User{
		RWMutex: new(sync.RWMutex),
	}

	if err := basic.getDatabase().GetConnection().Where("username = ?", username).Find(&user).Error; err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	if bcrypt.CompareHashAndPassword([]byte(*user.GetPassword()), []byte(password)) != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return user, nil
}

func (basic *Basic) Register(user *saas_model.User) (*saas_model.User, error) {
	bcrypted, err := bcrypt.GenerateFromPassword([]byte(*user.GetPassword()), 14)

	if err != nil {
		return nil, err
	}

	user.SetPassword(string(bcrypted))

	if err := basic.getDatabase().GetConnection().Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
