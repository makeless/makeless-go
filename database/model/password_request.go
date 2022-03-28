package makeless_go_model

import (
	"time"
)

type PasswordRequest struct {
	Model
	Email  string    `gorm:"not null"`
	Token  string    `gorm:"unique;not null"`
	Expire time.Time `gorm:"not null"`
	Used   bool      `gorm:"not null"`
}
