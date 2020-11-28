package web

type PollRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Participants []Participant `json:"participants"`
}

type Participant struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}

type PollResponse struct {
	ID int `json:"id"`
}
