package util

import "log"

type ErrorLogEvent struct {
	Function string
	Error error
	Params []LogEventParam
}

type InfoLogEvent struct {
	Function string
	Message string
	Params []LogEventParam
}

type LogEventParam struct {
	Key string
	Value string
}

func HandleFatal(event ErrorLogEvent) {
	if event.Error != nil {
		log.Fatal("FATAL ", event.Function, concatParams(event.Params), "details: ", event.Error)
	}
}

func HandleError(event ErrorLogEvent) {
	if event.Error != nil {
		log.Print("ERROR ", event.Function, concatParams(event.Params), "details: ", event.Error)
	}
}

func HandleInfo(event InfoLogEvent) {
	log.Print("INFO ", event.Function, ": ", formatMessage(event), concatParams(event.Params))
}

func concatParams(params []LogEventParam) string {
	result := " "
	if len(params) > 0 {
		result = ": "
		for _, param := range params {
			result += param.Key + " = '" + param.Value + "', "
		}
		result = result[:len(result)-2]
	}
	return result
}

func formatMessage(event InfoLogEvent) string {
	result := ""
	if len(event.Message) > 0 {
		result = event.Message
	}
	return result
}
