package _struct

import "sync"

type TeamInvitationTeamCreate struct {
	Invitations []*TeamInvitation `json:"invitations" binding:"min=1,max=5,dive"`
	*sync.RWMutex
}

func (teamInvitationTeamCreate *TeamInvitationTeamCreate) GetInvitations() []*TeamInvitation {
	teamInvitationTeamCreate.RLock()
	defer teamInvitationTeamCreate.RUnlock()

	return teamInvitationTeamCreate.Invitations
}
