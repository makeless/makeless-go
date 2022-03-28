package makeless_go_model

import "github.com/google/uuid"

type Team struct {
	Model
	Name string `gorm:"not null"`

	UserId uuid.UUID `gorm:"type:char(36);not null"`
	User   *User     `json:"user"`

	TeamUsers       []*TeamUser
	TeamInvitations []*TeamInvitation
}
