package saas_model

import "sync"

type Team struct {
	Model
	Name *string `gorm:"not null" json:"name" binding:"required,min=4"`

	Users []*User `gorm:"many2many:user_teams;" json:"-"`

	*sync.RWMutex `json:"-"`
}
