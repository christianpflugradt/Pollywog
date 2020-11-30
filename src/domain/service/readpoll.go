package service

import (
	"pollywog/db"
	"pollywog/domain/model"
)

func ReadPoll(secret string) (model.Poll, bool) {
	con := db.Database{}
	defer con.Disconnect()
	con.Connect()
	poll := con.SelectPoll(Hash(secret))
	if poll.ID != -1 {
		return poll, true
	} else {
		return poll, false
	}
}
