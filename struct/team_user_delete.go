package _struct

import "sync"

type TeamUserDelete struct {
	Id *uint `json:"id" binding:"required"`
	*sync.RWMutex
}

func (teamUserDelete *TeamUserDelete) GetId() *uint {
	teamUserDelete.RLock()
	defer teamUserDelete.RUnlock()

	return teamUserDelete.Id
}
