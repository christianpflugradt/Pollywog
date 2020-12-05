package db

func SetupTables(db Database) {
	setupTablePoll(db)
	setupTableParticipantInPoll(db)
	setupTableOptionInPoll(db)
	setupTableVoteInPoll(db)
}

func setupTablePoll(db Database) {
	db.executeDdl(`
		CREATE TABLE IF NOT EXISTS poll (
			id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(50) NOT NULL,
			description VARCHAR(255) NOT NULL,
			deadline DATE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
}

func setupTableParticipantInPoll(db Database) {
	db.executeDdl(`
		CREATE TABLE IF NOT EXISTS participant_in_poll (
			id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			poll_id INT UNSIGNED NOT NULL,
			displayname VARCHAR(50) NOT NULL,
			mail VARCHAR(50) NOT NULL,
			secret VARCHAR(64) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
}

func setupTableOptionInPoll(db Database) {
	db.executeDdl(`
		CREATE TABLE IF NOT EXISTS option_in_poll (
			id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			poll_id INT UNSIGNED NOT NULL,
			participant_id INT UNSIGNED NOT NULL,
			text VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
}

func setupTableVoteInPoll(db Database) {
	db.executeDdl(`
		CREATE TABLE IF NOT EXISTS vote_in_poll (
			id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			poll_id INT UNSIGNED NOT NULL,
			option_id INT UNSIGNED NOT NULL,
			participant_id INT UNSIGNED NOT NULL,
			weight INT UNSIGNED NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
}
