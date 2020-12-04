package transformer

import (
	"fmt"
	"pollywog/domain/model"
	"pollywog/web/representation"
	"time"
)

func pollToDomain(request representation.PollRequest) model.Poll {
	deadline, err := time.Parse("2006-01-02", request.Deadline)
	if err != nil {
		fmt.Print(err)
	}
	return model.Poll{
		Title:        request.Title,
		Description:  request.Description,
		Deadline:     deadline,
		Participants: participantsToDomain(request),
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
		Participants: participantsToRepresentation(poll),
		Options: optionsToRepresentation(poll),
	}
}
