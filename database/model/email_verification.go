package makeless_go_model

import "github.com/google/uuid"

type EmailVerification struct {
	Model
	Token    string `gorm:"unique;not null"`
	Verified bool   `gorm:"not null"`

	UserId uuid.UUID `gorm:"not null;type:uuid"`
	User   *User
}
