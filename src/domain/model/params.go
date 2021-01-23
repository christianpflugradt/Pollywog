package model

type PollParams struct {
	PollID        int
	OptionsPerParticipant int
	VotesPerParticipant int
}

type ParamKey string

const (
	OptionsPerParticipant ParamKey = "options-per-participant"
	VotesPerParticipant ParamKey = "votes-per-participant"
	PollCreator ParamKey = "poll-creator"
)
