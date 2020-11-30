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
}

type Participant struct {
	ID int
	Name string
	Mail string
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
