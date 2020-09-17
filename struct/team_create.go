package _struct

import "sync"

type TeamCreate struct {
	Name        *string           `json:"name" binding:"required,min=4,max=50"`
	Invitations []*TeamInvitation `json:"invitations" binding:"max=5,dive"`
	*sync.RWMutex
}

func (teamCreate *TeamCreate) GetName() *string {
	teamCreate.RLock()
	defer teamCreate.RUnlock()

	return teamCreate.Name
}

func (teamCreate *TeamCreate) GetInvitations() []*TeamInvitation {
	teamCreate.RLock()
	defer teamCreate.RUnlock()

	return teamCreate.Invitations
}
