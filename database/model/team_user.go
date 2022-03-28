package makeless_go_model

type TeamUser struct {
	Model
	TeamId uint  `basic:"not null"`
	Team   *Team `json:"team"`

	UserId uint  `basic:"not null"`
	User   *User `json:"user"`

	Role string `basic:"not null"`

	TeamInvitations []*TeamInvitation
}
