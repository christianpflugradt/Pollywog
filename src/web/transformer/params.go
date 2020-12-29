package transformer

import (
	"pollywog/domain/model"
	"pollywog/web/representation"
)

func paramsToDomain(request representation.PollRequest) model.PollParams {
	return model.PollParams{
		OptionsPerParticipant: request.Params.OptionsPerParticipant,
		VotesPerParticipant: request.Params.VotesPerParticipant,
	}
}

func paramsToRepresentation(poll model.Poll) representation.ParamsResponse {
	return representation.ParamsResponse{
		OptionsPerParticipant: poll.Params.OptionsPerParticipant,
		VotesPerParticipant: poll.Params.VotesPerParticipant,
	}
}
