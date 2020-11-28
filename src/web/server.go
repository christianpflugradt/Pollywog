package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pollywog/domain/service"
)

func postPoll(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var request PollRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err == nil {
			poll := toDomainObject(request)
			if service.IsValidForCreation(poll) {
				id := service.CreatePoll(poll)
				response := toPollResponse(id)
				json.NewEncoder(w).Encode(response)
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
			}
		} else {
			fmt.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func postOptions(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func postVotes(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func Serve() {
	http.HandleFunc("/poll", postPoll)
	http.HandleFunc("/options", postOptions)
	http.HandleFunc("/votes", postVotes)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
