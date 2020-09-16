package go_saas_model

import (
	"sync"
	"time"
)

type TeamInvitation struct {
	Model
	TeamId *uint `gorm:"not null" json:"teamId"`
	Team   *Team `json:"team"`

	UserId *uint `gorm:"not null" json:"userId"`
	User   *User `json:"user"`

	Email    *string    `gorm:"not null" json:"email"`
	Token    *string    `gorm:"unique;not null" json:"-"`
	Expire   *time.Time `gorm:"not null" json:"expire"`
	Accepted *bool      `gorm:"not null" json:"accepted"`

	*sync.RWMutex
}
