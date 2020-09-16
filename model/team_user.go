package go_saas_model

import (
	"sync"
)

type TeamUser struct {
	Model
	TeamId *uint `gorm:"not null" json:"teamId"`
	Team   *Team `json:"team"`

	UserId *uint `gorm:"not null" json:"userId"`
	User   *User `json:"user"`

	Role *string `gorm:"not null" json:"role"`

	*sync.RWMutex
}

func (teamUser *TeamUser) GetUser() *User {
	teamUser.RLock()
	defer teamUser.RUnlock()

	return teamUser.User
}
