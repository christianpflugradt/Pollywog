package service

import (
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"pollywog/db"
	"pollywog/domain/model"
	"pollywog/system"
	"strings"
	"time"
)

const secretSize = 64

func ResolveParticipant(secret string) (int, int) {
	con := db.Database{}
	defer con.Disconnect()
	con.Connect()
	return con.IdentifyParticipant(Hash(secret))
}

func supplySecrets(poll model.Poll, admintoken sys.Admintoken) {
	rand.Seed(time.Now().UnixNano())
	participants := poll.Participants
	for index, _ := range participants {
		unhashed := randomString()
		participants[index].Secret = Hash(unhashed)
		notifyParticipant(poll, admintoken, participants[index], unhashed)
	}
}

func Hash(secret string) string {
	hash := sha512.New512_256()
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum(nil))
}

func randomString() string {
	chars := []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	bytes := make([]rune, secretSize)
	for index := range bytes {
		bytes[index] = chars[rand.Intn(len(chars))]
	}
	return string(bytes)
}

func notifyParticipant(poll model.Poll, admintoken sys.Admintoken, participant model.Participant, unhashed string) {
	var config *sys.Config
	to := []string{participant.Mail}
	msg := []byte("To: " + participant.Mail +
		"\r\nContent-Type: text/plain; charset=\"UTF-8\";\r\n" +
		"Subject: invitation to poll: " + poll.Title + "\r\n\r\n" +
		"Hi " + participant.Name + "!\r\n" +
		"You are invited to participate in a poll.\r\n\r\n" +
		"Title: " + poll.Title + "\r\n" +
		"Description: " + poll.Description + "\r\n\r\n" +
		"Use the following link to participate: " + config.Get().Client.BaseUrl + "#" + unhashed + "\r\n\r\n" +
		"Best regards,\r\nPollywog \U0001F438" + invitedBy(admintoken.User))
	sys.SendMail(to, msg)
}

func invitedBy(user string) string {
	if len(user) > 0 {
		return "\r\n\r\n(invited by " + user + ")"
	} else {
		return ""
	}
}

func IsAdminAuthorizedToInviteParticipants(poll model.Poll, admintoken sys.Admintoken) string {
	if len(admintoken.Whitelist) > 0 {
		for _, participant := range poll.Participants {
			matches := false
			for _, whitelistEntry := range admintoken.Whitelist {
				if strings.HasSuffix(participant.Mail, whitelistEntry) {
					matches = true
					break
				}
			}
			if !matches {
				return "not allowed to send mails to " + participant.Mail
			}
		}
	}
	return ""
}

func IsVerifiedAdmin(secret string) (string, sys.Admintoken) {
	if secret == "" {
		return "no credentials provided", sys.Admintoken{}
	}
	var config *sys.Config
	if deprecatedIsVerifiedAdmin(secret) {
		return "", sys.Admintoken{}
	} else {
		for _, admintoken := range config.Get().Server.Admintokens {
			if admintoken.Token == Hash(secret) {
				return "", admintoken
			}
		}
		return "invalid credentials provided", sys.Admintoken{}
	}
}

func deprecatedIsVerifiedAdmin(secret string) bool {
	var config *sys.Config
	return Hash(config.Get().Server.Admintoken) == secret
}
