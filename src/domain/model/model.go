package model

import "time"

type Poll struct {
	ID int
	Title string
	Desc string
	Open bool
	Deadline time.Time
}

type Participant struct {
	ID int
	Name string
	Email string
	Secret string
}

type PollOption struct {
	ID int
	Title string
	PollID int
	ParticipantID int
}

type PollOptionVote struct {
	ID int
	PollOptionID int
	ParticipantID int
}
