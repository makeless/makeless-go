package makeless_go_model

type EmailVerification struct {
	Model
	Token    string `basic:"unique;not null"`
	Verified bool   `basic:"not null"`

	UserId uint `basic:"not null"`
	User   *User
}
