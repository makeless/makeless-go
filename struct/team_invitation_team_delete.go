package _struct

import "sync"

type TeamInvitationTeamDelete struct {
	Id *uint `json:"id" binding:"required"`
	*sync.RWMutex
}

func (teamInvitationTeamDelete *TeamInvitationTeamDelete) GetId() *uint {
	teamInvitationTeamDelete.RLock()
	defer teamInvitationTeamDelete.RUnlock()

	return teamInvitationTeamDelete.Id
}
