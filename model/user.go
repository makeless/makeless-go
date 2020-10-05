package makeless_go_model

import (
	"sync"
)

type User struct {
	Model
	Name              *string            `gorm:"not null" json:"name"`
	Password          *string            `gorm:"not null" json:"-"`
	Email             *string            `gorm:"unique;not null" json:"email"`
	EmailVerification *EmailVerification `gorm:"" json:"emailVerification"`

	TeamUsers       []*TeamUser       `json:"teamUsers" binding:"-"`
	TeamInvitations []*TeamInvitation `json:"teamInvitations" binding:"-"`
	Tokens          []*Token          `json:"tokens" binding:"-"`

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

func (user *User) GetEmailVerification() *EmailVerification {
	user.RLock()
	defer user.RUnlock()

	return user.EmailVerification
}

func (user *User) GetTeamUsers() []*TeamUser {
	user.RLock()
	defer user.RUnlock()

	return user.TeamUsers
}

func (user *User) GetTokens() []*Token {
	user.RLock()
	defer user.RUnlock()

	return user.Tokens
}
