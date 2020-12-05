package db

import (
	"fmt"
	"pollywog/domain/model"
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
		fmt.Print(err)
	}
	return model.Poll{
		ID: id,
		Title: title,
		Description: description,
		Deadline: deadline,
		Participants: db.selectPollParticipants(id),
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
		fmt.Print(err)
	}
	return model.Poll{
		ID: id,
		RequesterID: requesterId,
		Title: title,
		Description: description,
		Deadline: deadline,
		Participants: db.selectPollParticipants(id),
		Options: db.selectPollOptions(id),
		Votes: db.selectPollOptionVotes(id),
	}
}

func (db *Database) selectPollParticipants(id int) []model.Participant {
	rows, err := db.con.Query("SELECT id, displayname FROM participant_in_poll WHERE poll_id = ? ORDER BY id", id)
	if err != nil {
		fmt.Print(err)
	}
	participants := make([]model.Participant, 0)
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Print(err)
		}
		participants = append(participants, model.Participant{ ID: id, Name: name })
	}
	return participants
}

func (db *Database) selectPollOptions(id int) []model.PollOption {
	rows, err := db.con.Query("SELECT id, participant_id, text FROM option_in_poll WHERE poll_id = ? ORDER BY id", id)
	if err != nil {
		fmt.Print(err)
	}
	options := make([]model.PollOption, 0)
	for rows.Next() {
		var id, participantId int
		var text string
		err := rows.Scan(&id, &participantId, &text)
		if err != nil {
			fmt.Print(err)
		}
		options = append(options, model.PollOption { ID: id, ParticipantID: participantId, Text: text })
	}
	return options
}

func (db *Database) selectPollOptionVotes(id int) []model.PollOptionVote {
	rows, err := db.con.Query("SELECT option_id, participant_id, weight FROM vote_in_poll WHERE poll_id = ? ORDER BY id", id)
	if err != nil {
		fmt.Print(err)
	}
	votes := make([]model.PollOptionVote, 0)
	for rows.Next() {
		var optionId, participantId, weight int
		err := rows.Scan(&optionId, &participantId, &weight)
		if err != nil {
			fmt.Print(err)
		}
		votes = append(votes, model.PollOptionVote { PollOptionID: id, ParticipantID: participantId, Weight: weight })
	}
	return votes
}
