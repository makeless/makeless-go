package makeless_go_model

import "github.com/google/uuid"

type Team struct {
	Model
	Name string `gorm:"not null"`

	UserId uuid.UUID `gorm:"not null"`
	User   *User     `json:"user"`

	TeamUsers       []*TeamUser
	TeamInvitations []*TeamInvitation
}
