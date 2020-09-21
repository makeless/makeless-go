package _struct

import "sync"

type TeamInvitationDecline struct {
	Id     *uint   `json:"id" binding:"required"`
	Token  *string `json:"token" binding:"required"`
	TeamId *uint   `json:"teamId" binding:"required"`
	*sync.RWMutex
}

func (teamInvitationDecline *TeamInvitationDecline) GetId() *uint {
	teamInvitationDecline.RLock()
	defer teamInvitationDecline.RUnlock()

	return teamInvitationDecline.Id
}

func (teamInvitationDecline *TeamInvitationDecline) GetTeamId() *uint {
	teamInvitationDecline.RLock()
	defer teamInvitationDecline.RUnlock()

	return teamInvitationDecline.TeamId
}

func (teamInvitationDecline *TeamInvitationDecline) GetToken() *string {
	teamInvitationDecline.RLock()
	defer teamInvitationDecline.RUnlock()

	return teamInvitationDecline.Token
}
