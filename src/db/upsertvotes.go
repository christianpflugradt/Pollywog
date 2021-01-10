package db

import (
	"pollywog/domain/model"
	"pollywog/util"
)

func (db *Database) sqlDeleteObsoleteVotes(pollId int, participantId int, votes []model.PollOptionVote) {
	optionIds := make([]int, len(votes))
	for index, vote := range votes {
		optionIds[index] = vote.PollOptionID
	}
	inClause := ""
	if len(optionIds) > 0 {
		inClause = " AND option_id NOT IN (" + util.IntSliceToString(optionIds, ",") + ")"
	}
	_, err := db.con.Exec(`
		DELETE FROM vote_in_poll 
		WHERE poll_id = ? 
		AND participant_id = ? ` + inClause, pollId, participantId)
	util.HandleError(util.ErrorLogEvent{ Function: "db.sqlDeleteObsoleteVotes", Error: err })
}

func (db *Database) sqlInsertNewVotes(pollId int, participantId int, votes []model.PollOptionVote) {
	optionIdsVoted := db.selectOptionIdsFromVotes(pollId, participantId)
	votesToBeCreated := make([]model.PollOptionVote, 0)
	for _, vote := range votes {
		if !util.IntInSlice(optionIdsVoted, vote.PollOptionID) {
			votesToBeCreated = append(votesToBeCreated, vote)
		}
	}
	for _, vote := range votesToBeCreated {
		_, err := db.con.Exec(`INSERT INTO vote_in_poll 
				(poll_id, option_id, participant_id, weight) VALUES (?, ?, ?, 1)`,
			pollId, vote.PollOptionID, vote.ParticipantID)
		util.HandleError(util.ErrorLogEvent{ Function: "db.sqlInsertNewVotes", Error: err })
	}
}

func (db *Database) selectOptionIdsFromVotes(pollId int, participantId int) []int {
	rows, err := db.con.Query(`
		SELECT option_id FROM vote_in_poll 
		WHERE poll_id = ? 
		AND participant_id = ?`, pollId, participantId)
	util.HandleError(util.ErrorLogEvent{ Function: "db.selectOptionIdsFromVotes", Error: err })
	options := make([]int, 0)
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		util.HandleError(util.ErrorLogEvent{ Function: "db.selectOptionIdsFromVotes", Error: err })
		options = append(options, id)
	}
	return options
}

func (db *Database) sqlVerifyOptionsExist(pollId int, optionIds []int) bool {
	var count int
	inClause := "(" + util.IntSliceToString(optionIds, ",") + ")"
	err := db.con.QueryRow("SELECT COUNT(id) FROM option_in_poll WHERE poll_id = ? AND id IN " + inClause,
		pollId).Scan(&count)
	util.HandleError(util.ErrorLogEvent{ Function: "db.sqlVerifyOptionsExist", Error: err })
	return count == len(optionIds)
}
