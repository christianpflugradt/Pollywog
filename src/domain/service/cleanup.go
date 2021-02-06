package service

import (
	"pollywog/db"
	sys "pollywog/system"
	"pollywog/util"
	"time"
)

func ScheduleCleanup() {
	var config *sys.Config
	cleanupSettings := config.Get().Poll.Cleanup
	if cleanupSettings.Enabled {
		util.HandleInfo(util.InfoLogEvent{ Function: "service.ScheduleCleanup", Message: "poll cleanup enabled"})
		tick := time.NewTicker(time.Hour * time.Duration(cleanupSettings.IntervalInHours))
		go schedule(tick)
	}
}

func schedule(tick *time.Ticker) {
	for range tick.C {
		cleanupTask()
	}
}

func cleanupTask() {
	database := db.Database{}
	defer database.Disconnect()
	database.Connect()
	expiredPolls := database.SelectExpiredPolls()
	for _, expiredPoll := range expiredPolls {
		database.DeletePoll(expiredPoll)
	}
}
