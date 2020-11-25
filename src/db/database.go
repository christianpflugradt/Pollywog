package db

import (
	"database/sql"
	"fmt"
)

type Database struct {
	con *sql.DB
}

func (db *Database) Connect() {
	db.con = db.connectSqlite()
}

func (db *Database) connectSqlite() *sql.DB {
	con, err := sql.Open("sqlite3", "/home/cpf/pollywog.db")
	if err != nil {
		fmt.Print(err)
	}
	return con
}

func (db *Database) Disconnect() {
	db.con.Close()
	db.con = nil
}

func (db *Database) createTable(tableSql string) {
	db.con.Exec(tableSql)
}
