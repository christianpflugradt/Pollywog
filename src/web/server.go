package web

import (
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"pollywog/domain/service"
	sys "pollywog/system"
	"pollywog/util"
	"pollywog/web/representation"
	"pollywog/web/transformer"
	"syscall"
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
	errMsg, admintoken := service.IsVerifiedAdmin(r.Header.Get("Authorization"))
	if errMsg == "" {
		var request representation.PollRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err == nil {
			poll := transformer.TransformPollRequest(request)
			errMsg = service.IsAdminAuthorizedToInviteParticipants(poll, admintoken)
			if errMsg == "" {
				errMsg = service.IsValidForCreation(poll)
				if errMsg == "" {
					createdPoll := service.CreatePoll(poll, admintoken)
					util.HandleInfo(util.InfoLogEvent{ Function: "web.postPoll", Message: "poll created"})
					response := transformer.TransformDomainPoll(createdPoll)
					err = json.NewEncoder(w).Encode(response)
				} else {
					http.Error(w, errMsg, http.StatusUnprocessableEntity)
				}
			} else {
				http.Error(w, errMsg, http.StatusForbidden)
			}
		} else {
			http.Error(w, "unparseable request body: " + err.Error(), http.StatusBadRequest)
		}
		util.HandleError(util.ErrorLogEvent{ Function: "web.postPoll", Error: err })
	} else {
		http.Error(w, errMsg, http.StatusUnauthorized)
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
	listenForShutdownSignal()
	util.HandleInfo(util.InfoLogEvent{ Function: "web.Serve", Message: "starting Pollywog instance..." })
	http.HandleFunc("/poll", multiPoll)
	http.HandleFunc("/options", postOptions)
	http.HandleFunc("/votes", postVotes)
	var config *sys.Config
	err := http.ListenAndServe(":" + config.Get().Server.Port, nil)
	util.HandleFatal(util.ErrorLogEvent{ Function: "web.Serve", Error: err })
}

func listenForShutdownSignal() {
	shutdownHook := make(chan os.Signal)
	signal.Notify(shutdownHook, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-shutdownHook
		util.HandleInfo(util.InfoLogEvent{Function: "web.Serve", Message: "stopping Pollywog instance... (received SIGTERM)"})
		os.Exit(0)
	}()
}
