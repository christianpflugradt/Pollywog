package db

func setupDatabase(db Database) {
	db.Connect()
	db.createTable(`
		CREATE TABLE Poll 	
	`)
	db.Disconnect()
}
