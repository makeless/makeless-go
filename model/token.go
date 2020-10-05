package makeless_go_model

import (
	"sync"
)

type Token struct {
	Model
	Token *string `gorm:"unique;not null" json:"token"`
	Note  *string `gorm:"not null" json:"note"`

	UserId *uint `gorm:"not null" json:"userId"`
	User   *User `json:"user"`

	TeamId *uint `json:"teamId"`
	Team   *Team `json:"team"`

	*sync.RWMutex
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

func (token *Token) GetNote() *string {
	token.RLock()
	defer token.RUnlock()

	return token.Note
}

func (token *Token) GetUserId() *uint {
	token.RLock()
	defer token.RUnlock()

	return token.UserId
}

func (token *Token) GetUser() *User {
	token.RLock()
	defer token.RUnlock()

	return token.User
}

func (token *Token) GetTeamId() *uint {
	token.RLock()
	defer token.RUnlock()

	return token.TeamId
}

func (token *Token) GetTeam() *Team {
	token.RLock()
	defer token.RUnlock()

	return token.Team
}
