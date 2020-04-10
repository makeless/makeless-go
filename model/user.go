package saas_model

import (
	"sync"
)

type User struct {
	Model
	FirstName *string `gorm:"not null" json:"firstName" binding:"required"`
	LastName  *string `gorm:"not null" json:"lastName" binding:"required"`
	Username  *string `gorm:"unique;not null" json:"username" binding:"required"`
	Password  *string `gorm:"not null" json:"password,omitempty" binding:"required"`
	Email     *string `gorm:"unique;not null" json:"email" binding:"required"`

	Teams  []*Team  `gorm:"many2many:user_teams;" json:"teams"`
	Tokens []*Token `json:"tokens"`

	*sync.RWMutex `json:"-"`
}

func (user *User) GetFirstName() *string {
	user.RLock()
	defer user.RUnlock()

	return user.FirstName
}

func (user *User) GetLastName() *string {
	user.RLock()
	defer user.RUnlock()

	return user.LastName
}

func (user *User) GetUsername() *string {
	user.RLock()
	defer user.RUnlock()

	return user.Username
}

func (user *User) SetPassword(password string) {
	user.Lock()
	defer user.Unlock()

	user.Password = &password
}

func (user *User) GetPassword() *string {
	user.RLock()
	defer user.RUnlock()

	return user.Password
}

func (user *User) GetEmail() *string {
	user.RLock()
	defer user.RUnlock()

	return user.Email
}

func (user *User) GetTokens() []*Token {
	user.RLock()
	defer user.RUnlock()

	return user.Tokens
}
