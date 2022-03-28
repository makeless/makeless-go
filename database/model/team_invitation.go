package makeless_go_model

import (
	"time"
)

type TeamInvitation struct {
	Model
	TeamId uint `basic:"not null"`
	Team   *Team

	TeamUserId uint `basic:"not null"`
	TeamUser   *TeamUser

	Email    string    `basic:"not null"`
	Token    string    `basic:"unique;not null"`
	Expire   time.Time `basic:"not null"`
	Accepted bool      `basic:"not null"`
}
