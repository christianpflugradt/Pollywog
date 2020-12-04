package db

func SetupTablesSqlite(db Database) {
	setupTablePollSqlite(db)
	setupTableParticipantInPollSqlite(db)
	setupTableOptionInPollSqlite(db)
}

func setupTablePollSqlite(db Database) {
	db.executeDdl(`
		CREATE TABLE IF NOT EXISTS poll (
			id INTEGER NOT NULL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			deadline DATE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)	
	`)
}

func setupTableParticipantInPollSqlite(db Database) {
	db.executeDdl(`
		CREATE TABLE IF NOT EXISTS participant_in_poll (
			id INTEGER NOT NULL PRIMARY KEY,
			poll_id INTEGER NOT NULL,
			displayname TEXT NOT NULL,
			mail TEXT NOT NULL,
			secret TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
}

func setupTableOptionInPollSqlite(db Database) {
	db.executeDdl(`
		CREATE TABLE IF NOT EXISTS option_in_poll (
			id INTEGER NOT NULL PRIMARY KEY,
			poll_id INTEGER NOT NULL,
			participant_id INTEGER NOT NULL,
			text TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
}
