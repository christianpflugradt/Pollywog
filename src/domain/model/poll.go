package model

import "time"

type Poll struct {
	ID          int
	RequesterID int
	Title       string
	Description string
	Open        bool
	Deadline    time.Time
	Participants []Participant
	Options []PollOption
	Votes []PollOptionVote
}
