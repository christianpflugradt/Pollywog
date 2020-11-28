package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"pollywog/domain/model"
)

type Database struct {
	con *sql.DB
	IsSqlite3 bool
}

func (db *Database) Connect() {
	db.con = db.connectSqlite()
}

func (db *Database) connectSqlite() *sql.DB {
	db.IsSqlite3 = true
	con, err := sql.Open("sqlite3", "/tmp/pollywog.db")
	if err != nil {
		fmt.Print(err)
	}
	return con
}

func (db *Database) connectMySql() *sql.DB {
	con, err := sql.Open("mysql", "user:password@pollywog")
	if err != nil {
		fmt.Print(err)
	}
	return con
}

func (db *Database) Disconnect() {
	db.con.Close()
	db.con = nil
}

func (db *Database) executeDdl(tableSql string) {
	db.con.Exec(tableSql)
}

func (db *Database) InsertPoll(poll model.Poll) int {
	_, err := db.con.Exec("INSERT INTO Poll (title, description, deadline) VALUES (?, ?, ?)",
		poll.Title, poll.Desc, poll.Deadline)
	if err != nil {
		fmt.Print(err)
	}
	row := db.con.QueryRow("SELECT MAX(ID) FROM Poll")
	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Print(err)
	}
	return id
}
