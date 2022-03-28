package makeless_go_database_basic

import (
	"fmt"
	"gorm.io/gorm"
)

func (database *Database) GetConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		database.getUsername(),
		database.getPassword(),
		database.getHost(),
		database.getPort(),
		database.getDatabase(),
	)
}

func (database *Database) GetConnection() *gorm.DB {
	database.RLock()
	defer database.RUnlock()

	return database.connection
}

func (database *Database) SetConnection(connection *gorm.DB) {
	database.Lock()
	defer database.Unlock()

	database.connection = connection
}

func (database *Database) Connect(dialector gorm.Dialector) error {
	connection, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		return err
	}

	database.SetConnection(connection)
	return nil
}
