package web

type PollRequest struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Deadline string `json:"deadline"`
}

type PollResponse struct {
	ID int `json:"id"`
}
