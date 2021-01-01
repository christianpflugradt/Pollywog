package db

import (
	"pollywog/domain/model"
	"pollywog/util"
	"time"
)

func (db *Database) sqlSelectPollById(id int) model.Poll {
	var title, description string
	var deadline time.Time
	err := db.con.QueryRow(`
		SELECT p.title, p.description, p.deadline FROM poll p 
		WHERE p.id = ?`, id).Scan(&title, &description, &deadline)
	if err != nil {
		id = -1
	}
	util.HandleError(util.ErrorLogEvent{ Function: "db.selectPollById", Error: err })
	return model.Poll{
		ID: id,
		Title: title,
		Description: description,
		Deadline: deadline,
		Open: deadline.After(time.Now()),
		Participants: db.selectPollParticipants(id),
		Params: db.selectPollParams(id),
	}
}

func (db *Database) sqlSelectPoll(secret string) model.Poll {
	var id, requesterId int
	var title, description string
	var deadline time.Time
	err := db.con.QueryRow(`
		SELECT p.id, pip.id AS reqid, p.title, p.description, p.deadline FROM poll p 
		INNER JOIN participant_in_poll pip
		ON p.id = pip.poll_id
		WHERE pip.secret = ?
			`, secret).Scan(&id, &requesterId, &title, &description, &deadline)
	if err != nil {
		id = -1
	}
	util.HandleError(util.ErrorLogEvent{ Function: "db.selectPoll", Error: err })
	return model.Poll{
		ID: id,
		RequesterID: requesterId,
		Title: title,
		Description: description,
		Open: deadline.After(time.Now()),
		Deadline: deadline,
		Participants: db.selectPollParticipants(id),
		Options: db.selectPollOptions(id),
		Votes: db.selectPollOptionVotes(id),
		Params: db.selectPollParams(id),
	}
}

func (db *Database) selectPollParticipants(id int) []model.Participant {
	rows, err := db.con.Query("SELECT id, displayname FROM participant_in_poll WHERE poll_id = ? ORDER BY id", id)
	util.HandleError(util.ErrorLogEvent{ Function: "db.selectPollParticipants", Error: err })
	participants := make([]model.Participant, 0)
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		util.HandleError(util.ErrorLogEvent{ Function: "db.selectPollParticipants", Error: err })
		participants = append(participants, model.Participant{ ID: id, Name: name })
	}
	return participants
}

func (db *Database) selectPollOptions(id int) []model.PollOption {
	rows, err := db.con.Query("SELECT id, participant_id, text FROM option_in_poll WHERE poll_id = ? ORDER BY id", id)
	util.HandleError(util.ErrorLogEvent{ Function: "db.selectPollOptions", Error: err })
	options := make([]model.PollOption, 0)
	for rows.Next() {
		var id, participantId int
		var text string
		err = rows.Scan(&id, &participantId, &text)
		util.HandleError(util.ErrorLogEvent{ Function: "db.selectPollOptions", Error: err })
		options = append(options, model.PollOption { ID: id, ParticipantID: participantId, Text: text })
	}
	return options
}

func (db *Database) selectPollOptionVotes(id int) []model.PollOptionVote {
	rows, err := db.con.Query("SELECT option_id, participant_id, weight FROM vote_in_poll WHERE poll_id = ? ORDER BY id", id)
	util.HandleError(util.ErrorLogEvent{ Function: "db.selectPollOptionVotes", Error: err })
	votes := make([]model.PollOptionVote, 0)
	for rows.Next() {
		var optionId, participantId, weight int
		err = rows.Scan(&optionId, &participantId, &weight)
		util.HandleError(util.ErrorLogEvent{ Function: "db.selectPollOptionVotes", Error: err })
		votes = append(votes, model.PollOptionVote { PollOptionID: optionId, ParticipantID: participantId, Weight: weight })
	}
	return votes
}

func (db *Database) selectPollParams(id int) model.PollParams {
	return model.PollParams{
		OptionsPerParticipant: db.selectOptionsPerParticipant(id),
		VotesPerParticipant: db.selectVotesPerParticipant(id),
	}
}

func (db *Database) selectOptionsPerParticipant(id int) int {
	var optionsPerParticipant int
	err := db.con.QueryRow("SELECT paramvalue FROM poll_params WHERE poll_id = ? AND paramkey = ?",
		id, model.OptionsPerParticipant).Scan(&optionsPerParticipant)
	util.HandleError(util.ErrorLogEvent{ Function: "db.selectOptionsPerParticipant", Error: err })
	if err != nil {
		optionsPerParticipant = 999
	}
	return optionsPerParticipant
}

func (db *Database) selectVotesPerParticipant(id int) int {
	var votesPerParticipant int
	err := db.con.QueryRow("SELECT paramvalue FROM poll_params WHERE poll_id = ? AND paramkey = ?",
		id, model.VotesPerParticipant).Scan(&votesPerParticipant)
	util.HandleError(util.ErrorLogEvent{ Function: "db.selectVotesPerParticipant", Error: err })
	if err != nil {
		votesPerParticipant = 999
	}
	return votesPerParticipant
}
