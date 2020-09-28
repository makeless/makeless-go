package _struct

import "sync"

type TeamInvitationTeamResend struct {
	Id *uint `json:"id" binding:"required"`
	*sync.RWMutex
}

func (teamInvitationTeamResend *TeamInvitationTeamResend) GetId() *uint {
	teamInvitationTeamResend.RLock()
	defer teamInvitationTeamResend.RUnlock()

	return teamInvitationTeamResend.Id
}
