package go_saas_model

import (
	"sync"
	"time"
)

type PasswordRequest struct {
	Email  *string   `gorm:"not null" json:"email"`
	Token  *string   `gorm:"not null" json:"-"`
	Expire time.Time `gorm:"not null" json:"expire"`
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

func (passwordRequest PasswordRequest) GetExpire() time.Time {
	passwordRequest.RLock()
	defer passwordRequest.RUnlock()

	return passwordRequest.Expire
}
