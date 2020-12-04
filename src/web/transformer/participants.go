package transformer

import (
	"pollywog/domain/model"
	"pollywog/web/representation"
)

func participantsToDomain(request representation.PollRequest) []model.Participant {
	participants := make([]model.Participant, len(request.Participants))
	for index, participant := range request.Participants {
		participants[index] = model.Participant {
			Name: participant.Name,
			Mail: participant.Mail,
		}
	}
	return participants
}

func participantsToRepresentation(poll model.Poll) []representation.ParticipantResponse {
	participants := make([]representation.ParticipantResponse, len(poll.Participants))
	for index, participant := range poll.Participants {
		participants[index] = representation.ParticipantResponse {
			ID: participant.ID,
			Name: participant.Name,
		}
	}
	return participants
}
