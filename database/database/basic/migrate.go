package makeless_go_database_basic

func (database *Database) Migrate(dst ...interface{}) error {
	return database.GetConnection().AutoMigrate(dst)
}
