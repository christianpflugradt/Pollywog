package service

import (
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"pollywog/domain/model"
	"pollywog/system"
	"time"
)

const secretSize = 64

func supplySecrets(participants []model.Participant) {
	rand.Seed(time.Now().UnixNano())
	for index, _ := range participants {
		unhashed := randomString()
		participants[index].Secret = Hash(unhashed)
		notifyParticipant(participants[index], unhashed)
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

func notifyParticipant(participant model.Participant, unhashed string) {
	var config *sys.Config
	to := []string{participant.Mail}
	msg := []byte("To: " + participant.Mail +
		"\r\nSubject: invitation to poll\r\n\r\n" +
		config.Get().Client.BaseUrl + unhashed + "\r\n")
	sys.SendMail(to, msg)
}
