package _struct

import "sync"

type TokenDelete struct {
	Id *uint `json:"id" binding:"required"`
	*sync.RWMutex
}

func (tokenDelete *TokenDelete) GetId() *uint {
	tokenDelete.RLock()
	defer tokenDelete.RUnlock()

	return tokenDelete.Id
}
