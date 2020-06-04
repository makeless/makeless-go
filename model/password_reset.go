package go_saas_model

import "sync"

type PasswordReset struct {
	Password                *string `json:"password" binding:"required,min=3"`
	NewPassword             *string `json:"newPassword" binding:"required,min=3"`
	NewPasswordConfirmation *string `json:"newPasswordConfirmation" binding:"required,min=3,eqfield=NewPassword"`
	*sync.RWMutex
}

func (passwordReset *PasswordReset) GetPassword() *string {
	passwordReset.RLock()
	defer passwordReset.RUnlock()

	return passwordReset.Password
}

func (passwordReset *PasswordReset) GetNewPassword() *string {
	passwordReset.RLock()
	defer passwordReset.RUnlock()

	return passwordReset.NewPassword
}
