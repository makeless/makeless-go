package _struct

import "sync"

type TeamUserTeamDelete struct {
	Id uint `json:"id" binding:"required"`
	*sync.RWMutex
}

func (teamUserTeamDelete *TeamUserTeamDelete) GetId() uint {
	teamUserTeamDelete.RLock()
	defer teamUserTeamDelete.RUnlock()

	return teamUserTeamDelete.Id
}
