package _struct

import "sync"

type PasswordUpdate struct {
	Password                *string `json:"password" binding:"required,min=6"`
	NewPassword             *string `json:"newPassword" binding:"required,min=6"`
	NewPasswordConfirmation *string `json:"newPasswordConfirmation" binding:"required,min=6,eqfield=NewPassword"`
	*sync.RWMutex
}

func (passwordUpdate *PasswordUpdate) GetPassword() *string {
	passwordUpdate.RLock()
	defer passwordUpdate.RUnlock()

	return passwordUpdate.Password
}

func (passwordUpdate *PasswordUpdate) GetNewPassword() *string {
	passwordUpdate.RLock()
	defer passwordUpdate.RUnlock()

	return passwordUpdate.NewPassword
}

func (passwordUpdate *PasswordUpdate) GetNewPasswordConfirmation() *string {
	passwordUpdate.RLock()
	defer passwordUpdate.RUnlock()

	return passwordUpdate.NewPasswordConfirmation
}
