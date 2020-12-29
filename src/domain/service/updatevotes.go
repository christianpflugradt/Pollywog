package service

import (
	"pollywog/db"
	"pollywog/domain/model"
)

func UpdatePollOptionVotes(pollId int, votes []model.PollOptionVote) bool {
	if !isPollOpen(pollId) {
		return false
	}
	con := db.Database{}
	defer con.Disconnect()
	con.Connect()
	optionIds := make([]int, len(votes))
	for index, vote := range votes {
		optionIds[index] = vote.PollOptionID
	}
	participantId := votes[0].ParticipantID
	valid := con.VerifyOptionsExist(pollId, optionIds) && con.VerifyNumberOfVotesNotExceeded(pollId, votes)
	if valid {
		con.DeleteObsoleteVotes(pollId, participantId, votes)
		con.InsertNewVotes(pollId, participantId, votes)
	}
	return valid
}
