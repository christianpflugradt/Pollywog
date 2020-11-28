package service

import (
	"pollywog/db"
	"pollywog/domain/model"
	"time"
)

func IsValidForCreation(poll model.Poll) bool {
	valid := len(poll.Title) >= 3
	if valid {
		valid = isParticipantsValid(poll)
	}
	if valid {
		valid = isDeadlineValid(poll)
	}
	return valid
}

func isParticipantsValid(poll model.Poll) bool {
	valid := len(poll.Participants) > 1
	if valid {
		for _, participant := range poll.Participants {
			if len(participant.Name) <= 1 {
				valid = false
				break
			}
			if len(participant.Mail) <= 2 {
				valid = false
				break
			}
		}
	}
	return valid
}

func isDeadlineValid(poll model.Poll) bool {
	return poll.Deadline.Before(time.Now().AddDate(0, 3, 0)) &&
		poll.Deadline.After(time.Now().Add(time.Hour * time.Duration(1)))
}

func CreatePoll(poll model.Poll) int {
	supplySecrets(poll.Participants)
	con := db.Database{}
	defer con.Disconnect()
	con.Connect()
	id := con.InsertPoll(poll)
	return id
}
