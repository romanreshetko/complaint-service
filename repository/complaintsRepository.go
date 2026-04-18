package repository

import (
	"complaint-service/models"
	"database/sql"
	"errors"
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

func CreateCommentComplaint(db *sql.DB, req models.CreateCommentComplaint) error {
	_, err := db.Exec(`
		INSERT INTO comment_complaints (comment_id, author_id, reason, created_at)
		VALUES ($1, $2, $3, NOW())
`, req.CommentID, req.AuthorID, req.Reason)

	if err != nil {
		return err
	}

	return nil
}

func GetReviewComplaints(db *sql.DB) ([]models.ReviewComplaint, error) {
	rows, err := db.Query(`
		SELECT id, review_id, author_id, reason, created_at
		FROM review_complaints
`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var complaints []models.ReviewComplaint

	for rows.Next() {
		var complaint models.ReviewComplaint
		err := rows.Scan(
			&complaint.ID,
			&complaint.ReviewID,
			&complaint.AuthorID,
			&complaint.Reason,
			&complaint.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		complaints = append(complaints, complaint)
	}
	return complaints, nil
}

func GetCommentComplaints(db *sql.DB) ([]models.CommentComplaint, error) {
	rows, err := db.Query(`
		SELECT id, review_id, author_id, reason, created_at
		FROM comment_complaints
`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var complaints []models.CommentComplaint

	for rows.Next() {
		var complaint models.CommentComplaint
		err := rows.Scan(
			&complaint.ID,
			&complaint.CommentID,
			&complaint.AuthorID,
			&complaint.Reason,
			&complaint.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		complaints = append(complaints, complaint)
	}
	return complaints, nil
}

func DeleteReviewComplaints(db *sql.DB, reviewID int64) error {
	_, err := db.Exec(`
		DELETE FROM review_complaints 
		WHERE review_id = $1
`, reviewID)

	if err != nil {
		return err
	}
	return nil
}

func DeleteCommentComplaints(db *sql.DB, commentID int64) error {
	_, err := db.Exec(`
		DELETE FROM comment_complaints 
		WHERE comment_id = $1
`, commentID)

	if err != nil {
		return err
	}
	return nil
}

func GetReviewIDByComplaintID(db *sql.DB, complaintID int64) (int64, error) {
	var id int64
	err := db.QueryRow(`
		SELECT review_id
		FROM review_complaints
		WHERE id = $1
`, complaintID).Scan(&id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("complaint not found")
		}
		return 0, err
	}

	return id, nil
}

func GetCommentIDByComplaintID(db *sql.DB, complaintID int64) (int64, error) {
	var id int64
	err := db.QueryRow(`
		SELECT comment_id
		FROM comment_complaints
		WHERE id = $1
`, complaintID).Scan(&id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("complaint not found")
		}
		return 0, err
	}

	return id, nil
}
