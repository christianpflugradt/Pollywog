package service

import (
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"pollywog/db"
	"pollywog/domain/model"
	"pollywog/system"
	"time"
)

const secretSize = 64

func ResolveParticipant(secret string) (int, int) {
	con := db.Database{}
	defer con.Disconnect()
	con.Connect()
	return con.IdentifyParticipant(Hash(secret))
}

func supplySecrets(poll model.Poll) {
	rand.Seed(time.Now().UnixNano())
	participants := poll.Participants
	for index, _ := range participants {
		unhashed := randomString()
		participants[index].Secret = Hash(unhashed)
		notifyParticipant(poll, participants[index], unhashed)
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

func notifyParticipant(poll model.Poll, participant model.Participant, unhashed string) {
	var config *sys.Config
	to := []string{participant.Mail}
	msg := []byte("To: " + participant.Mail +
		"\r\nContent-Type: text/plain; charset=\"UTF-8\";\r\n" +
		"Subject: invitation to poll: " + poll.Title + "\r\n\r\n" +
		"Hi,漢ж\U0001F976 " + participant.Name + "!\r\n" +
		"you are invited to participate in a poll.\r\n\r\n" +
		"Title: " + poll.Title + "\r\n" +
		"Description: " + poll.Description + "\r\n\r\n" +
		"Use the following link to participate: " + config.Get().Client.BaseUrl + unhashed + "\r\n\r\n" +
		"Best regards,\r\nPollywog")
	sys.SendMail(to, msg)
}

func IsVerifiedAdmin(secret string) bool {
	var config *sys.Config
	return Hash(config.Get().Server.Admintoken) == secret
}
