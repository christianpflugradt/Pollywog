package service

import (
	"pollywog/db"
	"pollywog/domain/model"
)

func UpdatePollOptionVotes(pollId int, votes []model.PollOptionVote) bool {
	if len(votes) == 0 {
		return true
	}
	con := db.Database{}
	defer con.Disconnect()
	con.Connect()
	optionIds := make([]int, len(votes))
	for index, vote := range votes {
		optionIds[index] = vote.PollOptionID
	}
	participantId := votes[0].ParticipantID
	if con.VerifyOptionsExist(pollId, optionIds) {
		con.DeleteObsoleteVotes(pollId, participantId, votes)
		con.InsertNewVotes(pollId, participantId, votes)
		return true
	} else {
		return false
	}
}
