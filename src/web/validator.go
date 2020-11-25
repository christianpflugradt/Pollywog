package web

import "time"

func IsValidPollRequest(request PollRequest) bool {
	var valid = len(request.Title) >= 3
	if valid {
		var deadline, err = time.Parse("20060102", request.Deadline)
		if err != nil {
			valid = false
		} else {
			if deadline.After(time.Now().AddDate(0, 3, 0)) {
				valid = false
			}
			if deadline.Before(time.Now().Add(time.Hour * time.Duration(1))) {
				valid = false
			}
		}
	}
	return valid
}
