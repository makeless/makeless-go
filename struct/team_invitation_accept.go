package _struct

import "sync"

type TeamInvitationAccept struct {
	Id     *uint   `json:"id" binding:"required"`
	Token  *string `json:"token" binding:"required"`
	TeamId *uint   `json:"teamId" binding:"required"`
	*sync.RWMutex
}

func (teamInvitationAccept *TeamInvitationAccept) GetId() *uint {
	teamInvitationAccept.RLock()
	defer teamInvitationAccept.RUnlock()

	return teamInvitationAccept.Id
}

func (teamInvitationAccept *TeamInvitationAccept) GetTeamId() *uint {
	teamInvitationAccept.RLock()
	defer teamInvitationAccept.RUnlock()

	return teamInvitationAccept.TeamId
}

func (teamInvitationAccept *TeamInvitationAccept) GetToken() *string {
	teamInvitationAccept.RLock()
	defer teamInvitationAccept.RUnlock()

	return teamInvitationAccept.Token
}
