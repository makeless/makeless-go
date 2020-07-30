package _struct

import "sync"

type TokenCreate struct {
	Note  *string `json:"note" binding:"required,min=4,max=30"`
	Token *string `json:"token" binding:"required,len=32"`
	*sync.RWMutex
}

func (tokenCreate *TokenCreate) GetNote() *string {
	tokenCreate.RLock()
	defer tokenCreate.RUnlock()

	return tokenCreate.Note
}

func (tokenCreate *TokenCreate) GetToken() *string {
	tokenCreate.RLock()
	defer tokenCreate.RUnlock()

	return tokenCreate.Token
}
