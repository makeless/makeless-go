package go_saas_database

import "github.com/jinzhu/gorm"

type Database interface {
	GetConnection() *gorm.DB
	SetConnection(connection *gorm.DB)
	GetConnectionString() string
	Connect() error
	Migrate() error

	User
	Password
	PasswordRequest
	Team
	Member
	Profile
	Token
}
