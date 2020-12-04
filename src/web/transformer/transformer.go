package transformer

import (
	"pollywog/domain/model"
	"pollywog/web/representation"
)

func TransformPollRequest(request representation.PollRequest) model.Poll {
	return pollToDomain(request)
}

func TransformDomainPoll(poll model.Poll) representation.PollResponse {
	return pollToRepresentation(poll)
}

func TransformOptionsRequest(pollId int, participantId int, request representation.OptionsRequest) []model.PollOption {
	return optionsToDomain(pollId, participantId, request)
}
