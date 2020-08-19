package go_saas_database_basic

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type Database struct {
	connection *gorm.DB

	Dialect  string
	Username string
	Password string
	Database string
	Host     string
	Port     string
	*sync.RWMutex
}

func (database *Database) getDialect() string {
	database.RLock()
	defer database.RUnlock()

	return database.Dialect
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
