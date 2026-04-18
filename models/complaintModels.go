package models

import "time"

type AuthContext struct {
	UserID int64
	Role   string
}

type CreateReviewComplaint struct {
	ReviewID int64  `json:"review_id"`
	AuthorID int64  `json:"author_id"`
	Reason   string `json:"reason"`
}

type CreateCommentComplaint struct {
	CommentID int64  `json:"comment_id"`
	AuthorID  int64  `json:"author_id"`
	Reason    string `json:"reason"`
}

type ReviewComplaint struct {
	ID        int64     `json:"id"`
	ReviewID  int64     `json:"review_id"`
	AuthorID  int64     `json:"author_id"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentComplaint struct {
	ID        int64     `json:"id"`
	CommentID int64     `json:"review_id"`
	AuthorID  int64     `json:"author_id"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
}
