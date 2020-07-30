package _struct

import "sync"

type TokenTeamCreate struct {
	Note   *string `json:"note" binding:"required,min=4,max=30"`
	Token  *string `json:"token" binding:"required,len=32"`
	UserId *uint   `json:"userId" binding:"required"`
	*sync.RWMutex
}

func (tokenTeamCreate *TokenTeamCreate) GetNote() *string {
	tokenTeamCreate.RLock()
	defer tokenTeamCreate.RUnlock()

	return tokenTeamCreate.Note
}

func (tokenTeamCreate *TokenTeamCreate) GetToken() *string {
	tokenTeamCreate.RLock()
	defer tokenTeamCreate.RUnlock()

	return tokenTeamCreate.Token
}

func (tokenTeamCreate *TokenTeamCreate) GetUserId() *uint {
	tokenTeamCreate.RLock()
	defer tokenTeamCreate.RUnlock()

	return tokenTeamCreate.UserId
}
