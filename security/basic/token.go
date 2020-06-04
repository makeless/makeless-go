package saas_security_basic

import (
	"github.com/go-saas/go-saas/model"
	"sync"
)

func (basic *Basic) TokenLogin(token string) (*go_saas_model.User, error) {
	var user = &go_saas_model.User{
		RWMutex: new(sync.RWMutex),
	}

	err := basic.getDatabase().GetConnection().
		Preload("Tokens", "token = ?", token).
		Joins("JOIN tokens ON tokens.user_id=users.id AND tokens.token = ?", token).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	user.Tokens[0].RWMutex = new(sync.RWMutex)

	return user, nil
}
