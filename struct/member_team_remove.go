package _struct

import "sync"

type UserTeamDelete struct {
	Id uint `json:"id" binding:"required"`
	*sync.RWMutex
}

func (userTeamDelete *UserTeamDelete) GetId() uint {
	userTeamDelete.RLock()
	defer userTeamDelete.RUnlock()

	return userTeamDelete.Id
}
