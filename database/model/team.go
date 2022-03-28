package makeless_go_model

type Team struct {
	Model
	Name string `basic:"not null"`

	UserId uint  `basic:"not null"`
	User   *User `json:"user"`

	TeamUsers       []*TeamUser
	TeamInvitations []*TeamInvitation
}
