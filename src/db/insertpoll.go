package db

import (
	"pollywog/domain/model"
	"pollywog/util"
)

func (db *Database) sqlInsertPoll(poll model.Poll) int {
	_, err := db.con.Exec("INSERT INTO poll (title, description, deadline) VALUES (?, ?, ?)",
		poll.Title, poll.Description, poll.Deadline)
	util.HandleError(util.ErrorLogEvent{ Function: "db.sqlInsertPoll", Error: err })
	row := db.con.QueryRow("SELECT MAX(id) FROM poll")
	var id int
	err = row.Scan(&id)
	util.HandleError(util.ErrorLogEvent{ Function: "db.sqlInsertPoll", Error: err })
	db.insertPollParticipants(id, poll)
	db.insertPollParams(id, poll)
	return id
}

func (db *Database) insertPollParticipants(id int, poll model.Poll) {
	for _, participant := range poll.Participants {
		_, err := db.con.Exec(`INSERT INTO participant_in_poll 
				(poll_id, displayname, mail, secret) VALUES (?, ?, ?, ?)`,
			id, participant.Name, util.MaskMail(participant.Mail), participant.Secret)
		util.HandleError(util.ErrorLogEvent{ Function: "db.insertPollParticipants", Error: err })
	}
}

func (db *Database) insertPollParams(id int, poll model.Poll) {
	_, err := db.con.Exec("INSERT INTO poll_params (poll_id, paramkey, paramvalue) VALUES (?, ?, ?)",
		id, model.OptionsPerParticipant, poll.Params.OptionsPerParticipant)
	util.HandleError(util.ErrorLogEvent{ Function: "db.insertPollParams", Error: err })
	_, err = db.con.Exec("INSERT INTO poll_params (poll_id, paramkey, paramvalue) VALUES (?, ?, ?)",
		id, model.VotesPerParticipant, poll.Params.VotesPerParticipant)
	util.HandleError(util.ErrorLogEvent{ Function: "db.insertPollParams", Error: err })
}
