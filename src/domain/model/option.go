package model

type PollOption struct {
	New           bool
	ID            int
	Text          string
	PollID        int
	ParticipantID int
}
