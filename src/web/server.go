package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pollywog/domain/service"
	sys "pollywog/system"
)

func multiPoll(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postPoll(w, r)
	} else if r.Method == http.MethodGet {
		getPoll(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func postPoll(w http.ResponseWriter, r *http.Request) {
	var request PollRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err == nil {
		poll := toDomainObject(request)
		if service.IsValidForCreation(poll) {
			poll.ID = service.CreatePoll(poll)
			response := toPollResponse(poll)
			err = json.NewEncoder(w).Encode(response)
			if err != nil {
				fmt.Print(err)
			}
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}
	} else {
		fmt.Print(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getPoll(w http.ResponseWriter, r *http.Request) {
	poll, valid := service.ReadPoll(r.Header.Get("Authorization"))
	if valid {
		response := toPollResponse(poll)
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			fmt.Print(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
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
	http.HandleFunc("/poll", multiPoll)
	http.HandleFunc("/options", postOptions)
	http.HandleFunc("/votes", postVotes)
	var config *sys.Config
	log.Fatal(http.ListenAndServe(":" + config.Get().Server.Port, nil))
}
