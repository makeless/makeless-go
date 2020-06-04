package go_saas_basic_database

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type Database struct {
	connection *gorm.DB

	dialect  string
	username string
	password string
	database string
	host     string
	port     string
	*sync.RWMutex
}

func (database *Database) getDialect() string {
	database.RLock()
	defer database.RUnlock()

	return database.dialect
}

func (database *Database) getUsername() string {
	database.RLock()
	defer database.RUnlock()

	return database.username
}

func (database *Database) getPassword() string {
	database.RLock()
	defer database.RUnlock()

	return database.password
}

func (database *Database) getDatabase() string {
	database.RLock()
	defer database.RUnlock()

	return database.database
}

func (database *Database) getHost() string {
	database.RLock()
	defer database.RUnlock()

	return database.host
}

func (database *Database) getPort() string {
	database.RLock()
	defer database.RUnlock()

	return database.port
}
