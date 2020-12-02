package makeless_go_database

import "gorm.io/gorm"

type Database interface {
	GetConnection() *gorm.DB
	SetConnection(connection *gorm.DB)
	GetConnectionString() string
	Connect(dialector gorm.Dialector) error
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
