package go_saas_model

import (
	"sync"
)

type User struct {
	Model
	Name     *string `gorm:"not null" json:"name" binding:"required,min=4"`
	Username *string `json:"username"`
	Password *string `gorm:"not null" json:"password,omitempty" binding:"required"`
	Email    *string `gorm:"unique;not null" json:"email" binding:"required"`

	Teams  []*Team  `gorm:"many2many:user_teams;" json:"teams"`
	Tokens []*Token `json:"tokens"`

	*sync.RWMutex `json:"-"`
}

func (user *User) GetId() uint {
	return user.Id
}

func (user *User) GetName() *string {
	user.RLock()
	defer user.RUnlock()

	return user.Name
}

func (user *User) GetPassword() *string {
	user.RLock()
	defer user.RUnlock()

	return user.Password
}

func (user *User) SetPassword(password string) {
	user.Lock()
	defer user.Unlock()

	user.Password = &password
}

func (user *User) GetEmail() *string {
	user.RLock()
	defer user.RUnlock()

	return user.Email
}

func (user *User) GetTeams() []*Team {
	user.RLock()
	defer user.RUnlock()

	return user.Teams
}

func (user *User) GetTokens() []*Token {
	user.RLock()
	defer user.RUnlock()

	return user.Tokens
}
