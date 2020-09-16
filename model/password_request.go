package go_saas_model

import (
	"sync"
	"time"
)

type PasswordRequest struct {
	Model
	Email  *string    `gorm:"not null" json:"email"`
	Token  *string    `gorm:"unique;not null" json:"-"`
	Expire *time.Time `gorm:"not null" json:"expire"`
	Used   *bool      `gorm:"not null" json:"used"`

	*sync.RWMutex
}

func (passwordRequest PasswordRequest) GetEmail() *string {
	passwordRequest.RLock()
	defer passwordRequest.RUnlock()

	return passwordRequest.Email
}

func (passwordRequest PasswordRequest) GetToken() *string {
	passwordRequest.RLock()
	defer passwordRequest.RUnlock()

	return passwordRequest.Token
}

func (passwordRequest PasswordRequest) GetExpire() *time.Time {
	passwordRequest.RLock()
	defer passwordRequest.RUnlock()

	return passwordRequest.Expire
}
