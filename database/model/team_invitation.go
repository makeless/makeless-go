package makeless_go_model

import (
	"github.com/google/uuid"
	"time"
)

type TeamInvitation struct {
	Model
	TeamId uuid.UUID `gorm:"not null"`
	Team   *Team

	TeamUserId uuid.UUID `gorm:"not null"`
	TeamUser   *TeamUser

	Email    string    `gorm:"not null"`
	Token    string    `gorm:"unique;not null"`
	Expire   time.Time `gorm:"not null"`
	Accepted bool      `gorm:"not null"`
}
