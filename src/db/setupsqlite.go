package db

func SetupTablesSqlite(db Database) {
	setupTablePollSqlite(db)
	setupTableParticipantInPollSqlite(db)
	setupTableOptionInPollSqlite(db)
	setupTableVoteInPollSqlite(db)
	setupTablePollParamsSqlite(db)
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

func setupTableVoteInPollSqlite(db Database) {
	db.executeDdl(`
		CREATE TABLE IF NOT EXISTS vote_in_poll (
			id INTEGER NOT NULL PRIMARY KEY,
			poll_id INTEGER NOT NULL,
			option_id INTEGER NOT NULL,
			participant_id INTEGER NOT NULL,
			weight INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
}

func setupTablePollParamsSqlite(db Database) {
	db.executeDdl(`
		CREATE TABLE IF NOT EXISTS poll_params (
			id INTEGER NOT NULL PRIMARY KEY,
			poll_id INT UNSIGNED NOT NULL,
			paramkey TEXT NOT NULL,
			paramvalue TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
}
