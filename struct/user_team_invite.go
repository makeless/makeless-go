package _struct

import "sync"

type UserTeamInvite struct {
	Invitations []*TeamInvitation `json:"invitations" binding:"min=1,max=5,dive"`
	*sync.RWMutex
}

func (userTeamInvite *UserTeamInvite) GetInvitations() []*TeamInvitation {
	userTeamInvite.RLock()
	defer userTeamInvite.RUnlock()

	return userTeamInvite.Invitations
}
