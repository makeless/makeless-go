package _struct

import "sync"

type TeamInvitationDelete struct {
	Id *uint `json:"id" binding:"required"`
	*sync.RWMutex
}

func (teamInvitationDelete *TeamInvitationDelete) GetId() *uint {
	teamInvitationDelete.RLock()
	defer teamInvitationDelete.RUnlock()

	return teamInvitationDelete.Id
}
