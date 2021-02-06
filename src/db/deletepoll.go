package db

import (
	"pollywog/util"
	"strconv"
)

func (db *Database) sqlSelectExpiredPolls(statement string, days int) []int {
	ids := make([]int, 0)
	if len(statement) == 0 {
		statement = "SELECT id FROM poll WHERE deadline < CURRENT_DATE() - INTERVAL ? DAY"
	}
	rows, err := db.con.Query(statement, days)
	util.HandleError(util.ErrorLogEvent{ Function: "db.sqlSelectExpiredPolls", Error: err })
	if err == nil {
		for rows.Next() {
			var id int
			err = rows.Scan(&id)
			util.HandleError(util.ErrorLogEvent{Function: "db.sqlSelectExpiredPolls", Error: err})
			ids = append(ids, id)
		}
	}
	return ids
}

func (db *Database) sqlDeletePoll(id int) {
	params := make([]util.LogEventParam, 0)
	params = append(params, util.LogEventParam{Key: "poll_id", Value: strconv.Itoa(id)})
	util.HandleInfo(util.InfoLogEvent{ Function: "db.DeletePoll", Message: "poll deleted", Params: params })
	_, err := db.con.Exec("DELETE FROM participant_in_poll WHERE poll_id = ?", id)
	util.HandleError(util.ErrorLogEvent{ Function: "db.sqlDeletePoll", Error: err })
	_, err = db.con.Exec("DELETE FROM option_in_poll WHERE poll_id = ?", id)
	util.HandleError(util.ErrorLogEvent{ Function: "db.sqlDeletePoll", Error: err })
	_, err = db.con.Exec("DELETE FROM vote_in_poll WHERE poll_id = ?", id)
	util.HandleError(util.ErrorLogEvent{ Function: "db.sqlDeletePoll", Error: err })
	_, err = db.con.Exec("DELETE FROM poll_params WHERE poll_id = ?", id)
	util.HandleError(util.ErrorLogEvent{ Function: "db.sqlDeletePoll", Error: err })
	_, err = db.con.Exec("DELETE FROM poll WHERE id = ?", id)
	util.HandleError(util.ErrorLogEvent{ Function: "db.sqlDeletePoll", Error: err })
}
