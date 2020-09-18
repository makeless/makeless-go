package _struct

import "sync"

type UserTeamRemove struct {
	Id uint `json:"id" binding:"required"`
	*sync.RWMutex
}

func (userTeamRemove *UserTeamRemove) GetId() uint {
	userTeamRemove.RLock()
	defer userTeamRemove.RUnlock()

	return userTeamRemove.Id
}
