package service

import (
	"pollywog/db"
	"pollywog/domain/model"
)

func UpdatePollOptions(pollId int, participantId int, options []model.PollOption) bool {
	if !isPollOpen(pollId) {
		return false
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
	if con.VerifyParticipantOwnsOptions(participantId, optionIds) {
		con.UpdatePollOptions(pollId, participantId, options)
		return true
	} else {
		return false
	}
}
