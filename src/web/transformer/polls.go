package transformer

import (
	"pollywog/domain/model"
	"pollywog/util"
	"pollywog/web/representation"
	"time"
)

func pollToDomain(request representation.PollRequest) model.Poll {
	deadline, err := time.Parse("2006-01-02", request.Deadline)
	util.HandleError(util.ErrorLogEvent{ Function: "transformer.PollToDomain", Error: err })
	return model.Poll{
		Title:        request.Title,
		Description:  request.Description,
		Deadline:     deadline,
		Participants: participantsToDomain(request),
		Params: paramsToDomain(request),
	}
}

func pollToRepresentation(poll model.Poll) representation.PollResponse {
	deadline := poll.Deadline.Format("2006-01-02")
	return representation.PollResponse{
		Version: model.Version,
		ID: poll.ID,
		RequesterId: poll.RequesterID,
		Title: poll.Title,
		Description: poll.Description,
		Deadline: deadline,
		Open: poll.Open,
		Participants: participantsToRepresentation(poll),
		Options: optionsToRepresentation(poll),
		Votes: votesToRepresentation(poll),
		Params: paramsToRepresentation(poll),
	}
}
