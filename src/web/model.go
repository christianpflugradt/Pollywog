package web

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

type PollResponse struct {
	Version string `json:"version"`
	ID int                            `json:"id"`
	RequesterId int `json:"requester_id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Deadline    string                `json:"deadline"`
	Participants []ParticipantResponse `json:"participants"`
}

type ParticipantResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
}
