package representation

type PollRequest struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Deadline    string                `json:"deadline"`
	Participants []ParticipantRequest `json:"participants"`
	Params ParamsRequest			  `json:"params"`
}

type ParticipantRequest struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}

type ParamsRequest struct {
	OptionsPerParticipant int `json:"optionsPerParticipant"`
	VotesPerParticipant int `json:"votesPerParticipant"`
}
