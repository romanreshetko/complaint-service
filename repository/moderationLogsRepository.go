package repository

import (
	"database/sql"
	"log"
)

func InsertModerationLog(db *sql.DB, actorID, contentID int64, contentType, result string) {
	_, err := db.Exec(`
		INSERT INTO moderation_logs
		(actor_id, action_time, content_type, content_id, result)
		VALUES ($1, NOW(), $2, $3, $4)
`, actorID, contentID, contentType, result)

	if err != nil {
		log.Printf("failed to insert log %v", err)
		return
	}
}
