package _struct

import "sync"

type TeamInvitationAccept struct {
	Id *uint `json:"id" binding:"required"`
	*sync.RWMutex
}

func (teamInvitationAccept *TeamInvitationAccept) GetId() *uint {
	teamInvitationAccept.RLock()
	defer teamInvitationAccept.RUnlock()

	return teamInvitationAccept.Id
}
