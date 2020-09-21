package _struct

import "sync"

type TeamInvitationTeamDelete struct {
	Id    *uint   `json:"id" binding:"required"`
	Token *string `json:"token" binding:"required"`
	Email *string `json:"email" binding:"required"`
	*sync.RWMutex
}

func (teamInvitationTeamDelete *TeamInvitationTeamDelete) GetId() *uint {
	teamInvitationTeamDelete.RLock()
	defer teamInvitationTeamDelete.RUnlock()

	return teamInvitationTeamDelete.Id
}

func (teamInvitationTeamDelete *TeamInvitationTeamDelete) GetToken() *string {
	teamInvitationTeamDelete.RLock()
	defer teamInvitationTeamDelete.RUnlock()

	return teamInvitationTeamDelete.Token
}

func (teamInvitationTeamDelete *TeamInvitationTeamDelete) GetEmail() *string {
	teamInvitationTeamDelete.RLock()
	defer teamInvitationTeamDelete.RUnlock()

	return teamInvitationTeamDelete.Email
}
