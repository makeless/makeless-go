package go_saas_model

import "sync"

type Profile struct {
	Name          *string `gorm:"not null" json:"name" binding:"required,min=4"`
	*sync.RWMutex `json:"-"`
}

func (profile *Profile) GetName() *string {
	profile.RLock()
	defer profile.RUnlock()

	return profile.Name
}
