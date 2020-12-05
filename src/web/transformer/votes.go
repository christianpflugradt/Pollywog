package transformer

import (
	"pollywog/domain/model"
	"pollywog/web/representation"
)

func votesToDomain(participantId int, request representation.VotesRequest) []model.PollOptionVote {
	votes := make([]model.PollOptionVote, len(request.Votes))
	for index, vote := range request.Votes {
		votes[index] = model.PollOptionVote {
			PollOptionID: vote.OptionID,
			ParticipantID: participantId,
			Weight: 1,
		}
	}
	return votes
}

func votesToRepresentation(poll model.Poll) []representation.VoteResponse {
	votes := make([]representation.VoteResponse, len(poll.Votes))
	for index, vote := range poll.Votes {
		votes[index] = representation.VoteResponse {
			OptionID: vote.PollOptionID,
			ParticipantID: vote.ParticipantID,
			Weight: 1,
		}
	}
	return votes
}
