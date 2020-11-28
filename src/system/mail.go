package sys

import (
	"fmt"
	"net/smtp"
)

func SendMail(to []string, message []byte) {
	var config *Config
	auth := smtp.PlainAuth(
		config.Get().Smtp.Identity,
		config.Get().Smtp.User,
		config.Get().Smtp.Password,
		config.Get().Smtp.Host)
	err := smtp.SendMail(
		config.Get().Smtp.Host + ":" + config.Get().Smtp.Port,
		auth,
		config.Get().Smtp.User,
		to,
		message)
	if err != nil {
		fmt.Print(err)
	}
}
