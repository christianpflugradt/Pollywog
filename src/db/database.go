package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"pollywog/domain/model"
	sys "pollywog/system"
)

type Database struct {
	con *sql.DB
}

func (db *Database) Connect() {
	var config *sys.Config
	var err error
	db.con, err = sql.Open(config.Get().Database.Driver, config.Get().Database.DataSourceName)
	if err != nil {
		fmt.Print(err)
	}
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

func (db *Database) IdentifyParticipant(hashed string) (int, int) {
	var pollId, participantId int
	err := db.con.QueryRow("SELECT poll_id, id FROM participant_in_poll WHERE secret = ?",
		hashed).Scan(&pollId, &participantId)
	if err != nil {
		pollId = -1
		participantId = -1
		fmt.Print(err)
	}
	return pollId, participantId
}

func (db *Database) InsertPoll(poll model.Poll) int {
	return db.sqlInsertPoll(poll)
}

func (db *Database) SelectPoll(secret string) model.Poll {
	return db.sqlSelectPoll(secret)
}

func (db *Database) SelectPollById(id int) model.Poll {
	return db.sqlSelectPollById(id)
}

func (db *Database) VerifyParticipantOwnsOptions(participantId int, optionIds []int) bool {
	if len(optionIds) > 0 {
		return db.sqlVerifyParticipantOwnsOptions(participantId, optionIds)
	} else {
		return true
	}
}

func (db *Database) UpdatePollOptions(pollId int, participantId int, options []model.PollOption) {
	db.sqlUpsertPollOptions(pollId, participantId, options)
}

func (db *Database) VerifyOptionsExist(pollId int, optionIds []int) bool {
	if len(optionIds) > 0 {
		return db.sqlVerifyOptionsExist(pollId, optionIds)
	} else {
		return true
	}
}

func (db *Database) VerifyNumberOfOptionsNotExceeded(pollId int, options []model.PollOption) bool {
	return len(options) <= db.selectOptionsPerParticipant(pollId)
}

func (db *Database) VerifyNumberOfVotesNotExceeded(pollId int, votes []model.PollOptionVote) bool {
	count := 0
	for _, vote := range votes {
		count += vote.Weight
	}
	return count <= db.selectVotesPerParticipant(pollId)
}

func (db *Database) DeleteObsoleteVotes(pollId int, participantId int, votes []model.PollOptionVote) {
	db.sqlDeleteObsoleteVotes(pollId, participantId, votes)
}

func (db *Database) InsertNewVotes(pollId int, participantId int, votes []model.PollOptionVote) {
	db.sqlInsertNewVotes(pollId, participantId, votes)
}
