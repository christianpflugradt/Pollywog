package representation

type VotesRequest struct {
	Votes []Vote `json:"votes"`
}

type Vote struct {
	OptionID int `json:"option_id"`
	Weight int `json:"weight"`
}
