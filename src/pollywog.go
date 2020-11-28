package main

import (
	"pollywog/db"
	"pollywog/web"
)

func main() {
	prepareDatabase()
	web.Serve()
}

func prepareDatabase() {
	database := db.Database{}
	defer database.Disconnect()
	database.Connect()
	if database.IsSqlite3 {
		db.SetupTablesSqlite(database)
	} else {
		db.SetupTables(database)
	}
}
