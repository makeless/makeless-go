package makeless_go_model

type User struct {
	Model
	Name     string `basic:"not null"`
	Password string `basic:"not null"`
	Email    string `basic:"unique;not null"`

	EmailVerification *EmailVerification
	TeamUsers         []*TeamUser
}
