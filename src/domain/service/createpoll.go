package service

import (
	"pollywog/db"
	"pollywog/domain/model"
	sys "pollywog/system"
	"strconv"
	"strings"
	"time"
)

func IsValidForCreation(poll model.Poll) string {
	if len(strings.TrimSpace(poll.Title)) < 3 {
		return "poll title is too short (minimum 3 characters)"
	} else if len(poll.Title) > 50 {
		return "poll title is too long (maximum 50 characters allowed but " +
			strconv.Itoa(len(poll.Title)) + " characters received)"
	} else if len(poll.Description) > 255 {
		return "poll description is too long (maximum 255 characters alloweed but " +
			strconv.Itoa(len(poll.Description)) + " characters received)"
	}
	errMsg := isParticipantsValid(poll)
	if errMsg == "" {
		errMsg = isDeadlineValid(poll)
	}
	return errMsg
}

func isParticipantsValid(poll model.Poll) string {
	if len(poll.Participants) < 2 {
		return "not enough poll participants (minimum 2 participants)"
	}
	for _, participant := range poll.Participants {
		if len(participant.Name) <= 1 {
			return "participant name is too short (minimum 2 characters)"
		} else if len(participant.Name) > 50 {
			return "participant name is too long (maximum 50 characters allowed but " +
				strconv.Itoa(len(participant.Name)) + " characters received)"
		} else if len(participant.Mail) <= 2 {
			return "participant mail address is too short (minimum 3 characters)"
		} else if len(participant.Mail) > 50 {
			return "participant mail address is too long (maximum 50 characters allowed but " +
				strconv.Itoa(len(participant.Mail)) + " characters received)"
		}
	}
	return ""
}

func isDeadlineValid(poll model.Poll) string {
	if poll.Deadline.After(time.Now().AddDate(0, 3, 0)) {
		return "invalid deadline (must be at most 3 months in the future)"
	} else if poll.Deadline.Before(time.Now().Add(time.Hour * time.Duration(1))) {
		return "invalid deadline (must be at least 1 hour in the future)"
	} else {
		return ""
	}
}

func CreatePoll(poll model.Poll, admintoken sys.Admintoken) model.Poll {
	supplySecrets(poll, admintoken)
	con := db.Database{}
	defer con.Disconnect()
	con.Connect()
	id := con.InsertPoll(poll, admintoken)
	return con.SelectPollById(id)
}
