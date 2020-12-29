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
	valid := con.VerifyParticipantOwnsOptions(participantId, optionIds) &&
		con.VerifyNumberOfOptionsNotExceeded(pollId, options)
	if valid {
		con.UpdatePollOptions(pollId, participantId, options)
	}
	return valid
}
