package web

import (
	"fmt"
	"pollywog/domain/model"
	"time"
)

func toDomainObject(request PollRequest) model.Poll {
	deadline, err := time.Parse("20060102", request.Deadline)
	if err != nil {
		fmt.Print(err)
	}
	return model.Poll{
		Title: request.Title,
		Desc: request.Desc,
		Deadline: deadline,
	}
}

func toPollResponse(id int) PollResponse {
	return PollResponse{
		ID: id,
	}
}
