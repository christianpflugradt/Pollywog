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
		Title: request.Title,
		Description: request.Description,
		Deadline: deadline,
		Participants: mapParticipants(request),
	}
}

func mapParticipants(request PollRequest) []model.Participant {
	participants := make([]model.Participant, len(request.Participants))
	for index, participant := range request.Participants {
		participants[index] = model.Participant {
			Name: participant.Name,
			Mail: participant.Mail,
		}
	}
	return participants
}

func toPollResponse(id int) PollResponse {
	return PollResponse{
		ID: id,
	}
}
