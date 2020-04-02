package saas_database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/loeffel-io/go-saas/model"
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

func (database *Database) setConnection(connection *gorm.DB) {
	database.Lock()
	defer database.Unlock()

	database.connection = connection
}

func (database *Database) getConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		database.getUsername(),
		database.getPassword(),
		database.getHost(),
		database.getPort(),
		database.getDatabase(),
	)
}

func (database *Database) Connect() error {
	connection, err := gorm.Open(database.getDialect(), database.getConnectionString())

	if err != nil {
		return err
	}

	database.setConnection(connection)
	return nil
}

func (database *Database) GetConnection() *gorm.DB {
	database.RLock()
	defer database.RUnlock()

	return database.connection
}

func (database *Database) AutoMigrate() error {
	return database.GetConnection().AutoMigrate(
		new(saas_model.User),
		new(saas_model.Token),
	).Error
}
