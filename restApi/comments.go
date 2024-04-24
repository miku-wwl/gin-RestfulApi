package restapi

type CommentReq struct {
	MediaID     int    `json:"media_id"`
	UserID      int    `json:"user_id"`
	CommentText string `json:"comment_text"`
}
