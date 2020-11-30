package web

import (
	"fmt"
	"pollywog/domain/model"
	"time"
)

func toDomainObject(request PollRequest) model.Poll {
	deadline, err := time.Parse("20060102", request.Deadline)
	if err != nil {
		fmt.Print(err)
	}
	return model.Poll{
		Title:        request.Title,
		Description:  request.Description,
		Deadline:     deadline,
		Participants: mapParticipantRequests(request),
	}
}

func mapParticipantRequests(request PollRequest) []model.Participant {
	participants := make([]model.Participant, len(request.Participants))
	for index, participant := range request.Participants {
		participants[index] = model.Participant {
			Name: participant.Name,
			Mail: participant.Mail,
		}
	}
	return participants
}

func toPollResponse(poll model.Poll) PollResponse {
	deadline := poll.Deadline.Format("20060102")
	return PollResponse{
		Version: model.Version,
		ID: poll.ID,
		RequesterId: poll.RequesterID,
		Title: poll.Title,
		Description: poll.Description,
		Deadline: deadline,
		Participants: mapParticipantResponses(poll),
	}
}

func mapParticipantResponses(poll model.Poll) []ParticipantResponse {
	participants := make([]ParticipantResponse, len(poll.Participants))
	for index, participant := range poll.Participants {
		participants[index] = ParticipantResponse {
			ID: participant.ID,
			Name: participant.Name,
		}
	}
	return participants
}
