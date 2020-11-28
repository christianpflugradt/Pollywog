package db

func SetupTables(db Database) {
	setupTablePoll(db)
}

func setupTablePoll(db Database) {
	db.executeDdl(`
		CREATE TABLE IF NOT EXISTS Poll (
			id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(50) NOT NULL,
			description VARCHAR(255) NOT NULL,
			deadline DATE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
}
