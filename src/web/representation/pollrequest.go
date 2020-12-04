package representation

type PollRequest struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Deadline    string                `json:"deadline"`
	Participants []ParticipantRequest `json:"participants"`
}

type ParticipantRequest struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}
