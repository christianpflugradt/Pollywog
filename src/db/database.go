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
	err := db.con.Close()
	if err != nil {
		fmt.Print(err)
	}
	db.con = nil
}

func (db *Database) executeDdl(tableSql string) {
	_, err := db.con.Exec(tableSql)
	if err != nil {
		fmt.Print(err)
	}
}

func (db *Database) InsertPoll(poll model.Poll) int {
	return db.sqlInsertPoll(poll)
}

func (db *Database) SelectPoll(secret string) model.Poll {
	return db.sqlSelectPoll(secret)
}
