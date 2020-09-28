package _struct

import "sync"

type Register struct {
	Name                 *string `json:"name" binding:"required,min=4,max=50"`
	Email                *string `json:"email" binding:"required,email"`
	Password             *string `json:"password" binding:"required,min=6"`
	PasswordConfirmation *string `json:"passwordConfirmation" binding:"required,min=6,eqfield=Password"`
	LegalConfirmation    *bool   `json:"legalConfirmation" binding:"required"`
	*sync.RWMutex
}

func (register *Register) GetName() *string {
	register.RLock()
	defer register.RUnlock()

	return register.Name
}

func (register *Register) GetEmail() *string {
	register.RLock()
	defer register.RUnlock()

	return register.Email
}

func (register *Register) GetPassword() *string {
	register.RLock()
	defer register.RUnlock()

	return register.Password
}

func (register *Register) GetPasswordConfirmation() *string {
	register.RLock()
	defer register.RUnlock()

	return register.PasswordConfirmation
}

func (register *Register) GetLegalConfirmation() *bool {
	register.RLock()
	defer register.RUnlock()

	return register.LegalConfirmation
}
