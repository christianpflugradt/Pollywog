package main

import (
	"pollywog/db"
	"pollywog/domain/service"
	sys "pollywog/system"
	"pollywog/web"
)

func main() {
	prepareDatabase()
	service.ScheduleCleanup()
	web.Serve()
}

func prepareDatabase() {
	database := db.Database{}
	defer database.Disconnect()
	database.Connect()
	var config *sys.Config
	if config.Get().Database.Driver == "sqlite3" {
		db.SetupTablesSqlite(database)
	} else {
		db.SetupTables(database)
	}
}
