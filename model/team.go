package saas_model

import "sync"

type Team struct {
	Model
	Name *string `gorm:"not null" json:"name" binding:"required"`

	Users []*User `gorm:"many2many:user_teams;"`

	*sync.RWMutex `json:"-"`
}
