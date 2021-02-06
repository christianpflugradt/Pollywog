package db

import (
	"pollywog/domain/model"
	"pollywog/util"
)

func (db *Database) sqlUpsertPollOptions(pollId int, participantId int, options []model.PollOption) {
	db.deleteObsoleteOptions(pollId, participantId, options)
	db.createNewOptions(options)
}

func (db *Database) deleteObsoleteOptions(pollId int, participantId int, options []model.PollOption) {
	existingOptions := make([]int, 0)
	for _, item := range options {
		if !item.New {
			existingOptions = append(existingOptions, item.ID)
		}
	}
	inClause := ""
	if len(existingOptions) > 0 {
		inClause = " AND id NOT IN (" + util.IntSliceToString(existingOptions, ",") + ")"
	}
	_, err := db.con.Exec("DELETE FROM option_in_poll WHERE participant_id = ? " + inClause, participantId)
	util.HandleError(util.ErrorLogEvent{ Function: "db.deleteObsoleteOptions", Error: err })
	_, err = db.con.Exec(`
			DELETE FROM vote_in_poll
			WHERE poll_id = ?
			AND participant_id = ?
			AND NOT EXISTS
			(SELECT id FROM option_in_poll
			WHERE id = vote_in_poll.option_id)`, pollId, participantId)
	util.HandleError(util.ErrorLogEvent{ Function: "db.deleteObsoleteOptions", Error: err })
}

func (db *Database) createNewOptions(options []model.PollOption) {
	for _, option := range options {
		if option.New {
			optionText := option.Text
			if len(optionText) > 255 {
				optionText = optionText[:252] + "..."
			}
			_, err := db.con.Exec(`INSERT INTO option_in_poll 
				(poll_id, participant_id, text) VALUES (?, ?, ?)`,
				option.PollID, option.ParticipantID, optionText)
			util.HandleError(util.ErrorLogEvent{ Function: "db.createNewOptions", Error: err })
		}
	}
}

func (db *Database) sqlVerifyParticipantOwnsOptions(participantId int, optionIds []int) bool {
	var count int
	inClause := "(" + util.IntSliceToString(optionIds, ",") + ")"
	err := db.con.QueryRow("SELECT COUNT(id) FROM option_in_poll WHERE participant_id = ? AND id IN " + inClause,
		participantId).Scan(&count)
	util.HandleError(util.ErrorLogEvent{ Function: "db.sqlVerifyParticipantOwnsOptions", Error: err })
	return count == len(optionIds)
}
