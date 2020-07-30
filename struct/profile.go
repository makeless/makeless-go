package _struct

import "sync"

type Profile struct {
	Name *string `json:"name" binding:"required,min=4,max=50"`
	*sync.RWMutex
}

func (profile *Profile) GetName() *string {
	profile.RLock()
	defer profile.RUnlock()

	return profile.Name
}
