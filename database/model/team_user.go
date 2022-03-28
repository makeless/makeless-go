package makeless_go_model

import "github.com/google/uuid"

type TeamUser struct {
	Model
	TeamId uuid.UUID `gorm:"not null;type:uuid"`
	Team   *Team     `json:"team"`

	UserId uuid.UUID `gorm:"not null;type:uuid"`
	User   *User     `json:"user"`

	Role string `gorm:"not null"`

	TeamInvitations []*TeamInvitation
}
