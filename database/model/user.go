package makeless_go_model

type User struct {
	Model
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`

	EmailVerification *EmailVerification
	TeamUsers         []*TeamUser
}
