package go_saas_model

import (
	"sync"
)

type Token struct {
	Model
	Token *string `gorm:"unique;not null" json:"token" binding:"required,len=32"`
	Note  *string `gorm:"not null" json:"note" binding:"required,min=4,max=30"`

	UserId *uint `gorm:"not null" json:"userId" binding:"-"`
	User   *User `json:"user" binding:"-"`

	TeamId *uint `json:"teamId" binding:"-"`
	Team   *User `json:"team" binding:"-"`

	*sync.RWMutex `json:"-" binding:"-"`
}

func (token *Token) GetId() uint {
	token.RLock()
	defer token.RUnlock()

	return token.Id
}

func (token *Token) GetToken() *string {
	token.RLock()
	defer token.RUnlock()

	return token.Token
}

func (token *Token) GetUserId() *uint {
	token.RLock()
	defer token.RUnlock()

	return token.UserId
}

func (token *Token) GetTeamId() *uint {
	token.RLock()
	defer token.RUnlock()

	return token.TeamId
}
