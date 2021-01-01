package web

import (
	"encoding/json"
	"net/http"
	"pollywog/domain/service"
	sys "pollywog/system"
	"pollywog/util"
	"pollywog/web/representation"
	"pollywog/web/transformer"
)

func multiPoll(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postPoll(w, r)
	} else if r.Method == http.MethodGet {
		getPoll(w, r)
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func postPoll(w http.ResponseWriter, r *http.Request) {
	if service.IsVerifiedAdmin(r.Header.Get("Authorization")) {
		var request representation.PollRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err == nil {
			poll := transformer.TransformPollRequest(request)
			if service.IsValidForCreation(poll) {
				createdPoll := service.CreatePoll(poll)
				util.HandleInfo(util.InfoLogEvent{ Function: "web.postPoll", Message: "poll created"})
				response := transformer.TransformDomainPoll(createdPoll)
				err = json.NewEncoder(w).Encode(response)
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		util.HandleError(util.ErrorLogEvent{ Function: "web.postPoll", Error: err })
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func getPoll(w http.ResponseWriter, r *http.Request) {
	poll, valid := service.ReadPoll(r.Header.Get("Authorization"))
	if valid {
		response := transformer.TransformDomainPoll(poll)
		err := json.NewEncoder(w).Encode(response)
		util.HandleError(util.ErrorLogEvent{ Function: "web.getPoll", Error: err })
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func postOptions(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var request representation.OptionsRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err == nil {
			pollId, participantId := service.ResolveParticipant(r.Header.Get("Authorization"))
			if pollId != -1 {
				options := transformer.TransformOptionsRequest(pollId, participantId, request)
				valid := service.UpdatePollOptions(pollId, participantId, options)
				if valid {
					util.HandleInfo(util.InfoLogEvent{ Function: "web.postOptions", Message: "options updated"})
					getPoll(w, r)
				} else {
					w.WriteHeader(http.StatusUnprocessableEntity)
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		util.HandleError(util.ErrorLogEvent{ Function: "web.postOptions", Error: err })
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func postVotes(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var request representation.VotesRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err == nil {
			pollId, participantId := service.ResolveParticipant(r.Header.Get("Authorization"))
			if pollId != -1 {
				votes := transformer.TransformVotesRequest(participantId, request)
				valid := service.UpdatePollOptionVotes(pollId, participantId, votes)
				if valid {
					util.HandleInfo(util.InfoLogEvent{ Function: "web.postVotes", Message: "votes updated"})
					getPoll(w, r)
				} else {
					w.WriteHeader(http.StatusUnprocessableEntity)
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		util.HandleError(util.ErrorLogEvent{ Function: "web.postVotes", Error: err })
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func Serve() {
	http.HandleFunc("/poll", multiPoll)
	http.HandleFunc("/options", postOptions)
	http.HandleFunc("/votes", postVotes)
	var config *sys.Config
	err := http.ListenAndServe(":" + config.Get().Server.Port, nil)
	util.HandleFatal(util.ErrorLogEvent{ Function: "web.Serve", Error: err })
}
