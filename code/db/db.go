package db

// Init makes sure the database has the correct tables.
func Init() {

	InitSQLite()

	DB.AutoMigrate(&Contact{})
	DB.AutoMigrate(&Note{})
}

// Close the database.
func Close() {
	DB.Close()
}
