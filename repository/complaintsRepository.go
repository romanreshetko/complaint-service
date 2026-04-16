package repository

import (
	"complaint-service/models"
	"database/sql"
)

func CreateReviewComplaint(db *sql.DB, req models.CreateReviewComplaint) error {
	_, err := db.Exec(`
		INSERT INTO review_complaints (review_id, author_id, reason, created_at)
		VALUES ($1, $2, $3, NOW())
`, req.ReviewID, req.AuthorID, req.Reason)

	if err != nil {
		return err
	}

	return nil
}
