package _struct

import (
	"sync"
)

type Login struct {
	Email    *string `json:"email" binding:"required"`
	Password *string `json:"password" binding:"required"`

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
