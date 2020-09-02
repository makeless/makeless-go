package _struct

import "sync"

type PasswordRequest struct {
	Email *string `json:"email" binding:"required,email"`
	*sync.RWMutex
}

func (passwordRequest *PasswordRequest) GetEmail() *string {
	passwordRequest.RLock()
	defer passwordRequest.RUnlock()

	return passwordRequest.Email
}
