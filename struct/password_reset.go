package _struct

import "sync"

type PasswordReset struct {
	Password             *string `json:"password" binding:"required,min=6"`
	PasswordConfirmation *string `json:"passwordConfirmation" binding:"required,min=6,eqfield=Password"`
	*sync.RWMutex
}

func (passwordReset *PasswordReset) GetPassword() *string {
	passwordReset.RLock()
	defer passwordReset.RUnlock()

	return passwordReset.Password
}

func (passwordReset *PasswordReset) GetPasswordConfirmation() *string {
	passwordReset.RLock()
	defer passwordReset.RUnlock()

	return passwordReset.PasswordConfirmation
}
