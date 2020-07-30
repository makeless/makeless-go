package _struct

import "sync"

type Register struct {
	Name     *string `json:"name" binding:"required,min=4,max=50"`
	Password *string `gorm:"not null" json:"password" binding:"required,min=6"`
	Email    *string `gorm:"unique;not null" json:"email" binding:"required,email"`
	*sync.RWMutex
}

func (register *Register) GetName() *string {
	register.RLock()
	defer register.RUnlock()

	return register.Name
}

func (register *Register) GetPassword() *string {
	register.RLock()
	defer register.RUnlock()

	return register.Password
}

func (register *Register) GetEmail() *string {
	register.RLock()
	defer register.RUnlock()

	return register.Email
}
