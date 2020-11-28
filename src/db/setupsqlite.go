package db

func SetupTablesSqlite(db Database) {
	setupTablePollSqlite(db)
}

func setupTablePollSqlite(db Database) {
	db.executeDdl(`
		CREATE TABLE IF NOT EXISTS Poll (
			id INTEGER NOT NULL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			deadline DATE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)	
	`)
}
