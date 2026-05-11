package repository

import (
	"database/sql"
	"log"
)

func InsertModerationLog(db *sql.DB, actorID, contentID int64, contentType, result string) error {
	_, err := db.Exec(`
		INSERT INTO moderation_logs
		(actor_id, action_time, content_type, content_id, result)
		VALUES ($1, NOW(), $2, $3, $4)
`, actorID, contentType, contentID, result)

	if err != nil {
		log.Printf("failed to insert log %v", err)
		return err
	}
	return nil
}

func GetModerationErrorsNumber(db *sql.DB) (int64, error) {
	var count int64
	err := db.QueryRow(`
		SELECT COUNT(*) FROM moderation_logs
		WHERE result = 'moderation_error'
		AND action_time > NOW() - INTERVAL '24 hours'
`).Scan(&count)
	if err != nil {
		return -1, err
	}

	return count, nil
}
