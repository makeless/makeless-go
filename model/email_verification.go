package go_saas_model

import "sync"

type EmailVerification struct {
	Model
	Token    *string `gorm:"unique;not null" json:"-"`
	Verified *bool   `gorm:"not null" json:"verified"`

	UserId *uint `gorm:"not null" json:"userId"`
	User   *User `json:"user"`

	*sync.RWMutex
}

func (emailVerification *EmailVerification) GetId() uint {
	emailVerification.RLock()
	defer emailVerification.RUnlock()

	return emailVerification.Id
}

func (emailVerification *EmailVerification) GetToken() *string {
	emailVerification.RLock()
	defer emailVerification.RUnlock()

	return emailVerification.Token
}

func (emailVerification *EmailVerification) GetVerified() *bool {
	emailVerification.RLock()
	defer emailVerification.RUnlock()

	return emailVerification.Verified
}

func (emailVerification *EmailVerification) GetUser() *User {
	emailVerification.RLock()
	defer emailVerification.RUnlock()

	return emailVerification.User
}
