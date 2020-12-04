package service

import (
	"fmt"
	"pollywog/db"
	"pollywog/domain/model"
)

func UpdatePollOptions(participantId int, options []model.PollOption) bool {
	if len(options) == 0 {
		return true
	}
	con := db.Database{}
	defer con.Disconnect()
	con.Connect()
	optionIds := make([]int, 0)
	for _, option := range options {
		if !option.New {
			optionIds = append(optionIds, option.ID)
		}
	}
	fmt.Print("a")
	if con.VerifyParticipantOwnsOptions(participantId, optionIds) {
		fmt.Print("b")
		con.UpdatePollOptions(participantId, options)
		return true
	} else {
		fmt.Print("c")
		return false
	}
}
