package representation

type PollResponse struct {
	Version string `json:"version"`
	ID int                            `json:"id"`
	RequesterId int `json:"requester_id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Deadline    string                `json:"deadline"`
	Participants []ParticipantResponse `json:"participants"`
	Options []OptionResponse `json:"options"`
}

type ParticipantResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type OptionResponse struct {
	ID int `json:"id"`
	ParticipantID int `json:"participant_id"`
	Text string `json:"text"`
}
