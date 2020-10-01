package _struct

import "sync"

type TokenTeamDelete struct {
	Id *uint `json:"id" binding:"required"`
	*sync.RWMutex
}

func (tokenTeamDelete *TokenTeamDelete) GetId() *uint {
	tokenTeamDelete.RLock()
	defer tokenTeamDelete.RUnlock()

	return tokenTeamDelete.Id
}
