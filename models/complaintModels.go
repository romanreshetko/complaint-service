package models

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
