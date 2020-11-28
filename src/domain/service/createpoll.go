package service

import (
	"pollywog/db"
	"pollywog/domain/model"
	"time"
)

func IsValidForCreation(poll model.Poll) bool {
	var valid = len(poll.Title) >= 3
	if valid {
		if poll.Deadline.After(time.Now().AddDate(0, 3, 0)) {
			valid = false
		}
		if poll.Deadline.Before(time.Now().Add(time.Hour * time.Duration(1))) {
			valid = false
		}
	}
	return valid
}

func CreatePoll(poll model.Poll) int {
	con := db.Database{}
	defer con.Disconnect()
	con.Connect()
	id := con.InsertPoll(poll)
	return id
}
