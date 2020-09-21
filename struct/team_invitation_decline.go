package _struct

import "sync"

type TeamInvitationDelete struct {
	Id     *uint   `json:"id" binding:"required"`
	Token  *string `json:"token" binding:"required"`
	TeamId *uint   `json:"teamId" binding:"required"`
	*sync.RWMutex
}

func (teamInvitationDelete *TeamInvitationDelete) GetId() *uint {
	teamInvitationDelete.RLock()
	defer teamInvitationDelete.RUnlock()

	return teamInvitationDelete.Id
}

func (teamInvitationDelete *TeamInvitationDelete) GetTeamId() *uint {
	teamInvitationDelete.RLock()
	defer teamInvitationDelete.RUnlock()

	return teamInvitationDelete.TeamId
}

func (teamInvitationDelete *TeamInvitationDelete) GetToken() *string {
	teamInvitationDelete.RLock()
	defer teamInvitationDelete.RUnlock()

	return teamInvitationDelete.Token
}
