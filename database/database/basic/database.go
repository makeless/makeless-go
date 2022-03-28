package makeless_go_database_basic

import (
	"gorm.io/gorm"
	"sync"
)

type Database struct {
	connection *gorm.DB

	Username string
	Password string
	Database string
	Host     string
	Port     string
	*sync.RWMutex
}

func (database *Database) getUsername() string {
	database.RLock()
	defer database.RUnlock()

	return database.Username
}

func (database *Database) getPassword() string {
	database.RLock()
	defer database.RUnlock()

	return database.Password
}

func (database *Database) getDatabase() string {
	database.RLock()
	defer database.RUnlock()

	return database.Database
}

func (database *Database) getHost() string {
	database.RLock()
	defer database.RUnlock()

	return database.Host
}

func (database *Database) getPort() string {
	database.RLock()
	defer database.RUnlock()

	return database.Port
}
