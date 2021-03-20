package sys

import (
	"net/smtp"
	"pollywog/util"
)

func SendMail(to []string, message []byte) {
	var config *Config
	var auth smtp.Auth
	if len(config.Get().Smtp.Password) > 0 {
		auth = smtp.PlainAuth(
			config.Get().Smtp.Identity,
			config.Get().Smtp.User,
			config.Get().Smtp.Password,
			config.Get().Smtp.Host)
	}
	err := smtp.SendMail(
		config.Get().Smtp.Host + ":" + config.Get().Smtp.Port,
		auth,
		config.Get().Smtp.User,
		to,
		message)
	util.HandleError(util.ErrorLogEvent{ Function: "sys.SendMail", Error: err })
}
