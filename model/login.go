package go_saas_model

import (
	"sync"
)

type Login struct {
	Password *string `binding:"required"`
	Email    *string `binding:"required"`

	*sync.RWMutex `json:"-"`
}

func (login *Login) GetPassword() *string {
	login.RLock()
	defer login.RUnlock()

	return login.Password
}

func (login *Login) GetEmail() *string {
	login.RLock()
	defer login.RUnlock()

	return login.Email
}
