package makeless_go_database

import "github.com/jinzhu/gorm"

type Database interface {
	GetConnection() *gorm.DB
	SetConnection(connection *gorm.DB)
	GetConnectionString() string
	Connect() error
	Migrate() error

	User
	EmailVerification
	Password
	PasswordRequest
	Team
	Profile
	Token
	TeamInvitation
}
