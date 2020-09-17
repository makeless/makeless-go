package _struct

import "sync"

type TeamInvitation struct {
	Email *string `json:"email" binding:"required,email"`
	*sync.RWMutex
}

func (teamInvitation *TeamInvitation) GetEmail() *string {
	teamInvitation.RLock()
	defer teamInvitation.RUnlock()

	return teamInvitation.Email
}
