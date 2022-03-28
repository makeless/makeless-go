package makeless_go_model

import (
	"time"
)

type PasswordRequest struct {
	Model
	Email  string    `basic:"not null"`
	Token  string    `basic:"unique;not null"`
	Expire time.Time `basic:"not null"`
	Used   bool      `basic:"not null"`
}
