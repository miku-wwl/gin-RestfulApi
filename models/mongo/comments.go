package mongodb

import (
	"context"
	"restApi/dao"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Comment struct {
	CommentID   int       `json:"comment_id"`
	MediaID     int       `json:"media_id"`
	UserID      int       `json:"user_id"`
	CommentText string    `json:"comment_text"`
	CreatedAt   time.Time `json:"created_at"`
}

func CreateComment(c Comment) (*mongo.InsertOneResult, error) {
	ret, err := dao.MonCollection.InsertOne(context.Background(), c)
	return ret, err
}
