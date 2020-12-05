package db

import (
	"fmt"
	"pollywog/domain/model"
)

func (db *Database) sqlInsertPoll(poll model.Poll) int {
	_, err := db.con.Exec("INSERT INTO poll (title, description, deadline) VALUES (?, ?, ?)",
		poll.Title, poll.Description, poll.Deadline)
	if err != nil {
		fmt.Print(err)
	}
	row := db.con.QueryRow("SELECT MAX(id) FROM poll")
	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Print(err)
	}
	db.insertPollParticipants(id, poll)
	return id
}

func (db *Database) insertPollParticipants(id int, poll model.Poll) {
	for _, participant := range poll.Participants {
		_, err := db.con.Exec(`INSERT INTO participant_in_poll 
				(poll_id, displayname, mail, secret) VALUES (?, ?, ?, ?)`,
			id, participant.Name, participant.Mail, participant.Secret)
		if err != nil {
			fmt.Print(err)
		}
	}
}
