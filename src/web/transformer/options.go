package transformer

import (
	"pollywog/domain/model"
	"pollywog/web/representation"
)

func optionsToDomain(pollId int, participantId int, request representation.OptionsRequest) []model.PollOption {
	optionsCount := len(request.CreateOptions) + len(request.KeepOptions)
	options := make([]model.PollOption, optionsCount)
	for index, id := range request.KeepOptions {
		options[index] = model.PollOption {
			PollID: pollId,
			ParticipantID: participantId,
			New: false,
			ID: id,
		}
	}
	offset := len(request.KeepOptions)
	for index, text := range request.CreateOptions {
		options[offset + index] = model.PollOption {
			PollID: pollId,
			ParticipantID: participantId,
			New: true,
			Text: text,
		}
	}
	return options
}

func optionsToRepresentation(poll model.Poll) []representation.OptionResponse {
	options := make([]representation.OptionResponse, len(poll.Options))
	for index, option := range poll.Options {
		options[index] = representation.OptionResponse {
			ID: option.ID,
			ParticipantID: option.ParticipantID,
			Text: option.Text,
		}
	}
	return options
}
