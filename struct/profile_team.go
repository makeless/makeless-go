package _struct

import "sync"

type ProfileTeam struct {
	Name *string `json:"name" binding:"required,min=4,max=50"`
	*sync.RWMutex
}

func (profileTeam *ProfileTeam) GetName() *string {
	profileTeam.RLock()
	defer profileTeam.RUnlock()

	return profileTeam.Name
}
