package models

type CreateLogRequest struct {
	ContentType string `json:"content_type"`
	ContentID   int64  `json:"content_id"`
	Result      string `json:"result"`
}
