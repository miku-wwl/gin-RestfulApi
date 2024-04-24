package mongodb

import (
	"context"
	"restApi/dao"
	"restApi/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Comment struct {
	CommentID   string    `json:"comment_id" bson:"comment_id"`
	MediaID     int       `json:"media_id" bson:"media_id"`
	UserID      int       `json:"user_id" bson:"user_id"`
	CommentText string    `json:"comment_text" bson:"comment_text"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}

func CreateComment(c *Comment) (*mongo.InsertOneResult, error) {
	ret, err := dao.MonCollection.InsertOne(context.Background(), *c)

	c.CommentID = ret.InsertedID.(primitive.ObjectID).Hex()

	updateIDFilter := bson.M{"_id": ret.InsertedID}
	updateCommentID := bson.M{"$set": bson.M{"comment_id": c.CommentID}}
	_, updateErr := dao.MonCollection.UpdateOne(context.Background(), updateIDFilter, updateCommentID)
	if updateErr != nil {
		logger.Error(map[string]interface{}{"[CreateComment] dao.MonCollection.UpdateOne error:": updateErr.Error()})
	}
	return ret, err
}

func GetCommentList() (*mongo.Cursor, error) {
	cur, err := dao.MonCollection.Find(context.Background(), bson.M{})
	return cur, err
}

func UpdateComment(c Comment) (*mongo.UpdateResult, error) {
	filter := bson.M{"comment_id": c.CommentID}
	update := bson.M{"$set": c}

	ret, err := dao.MonCollection.UpdateOne(context.Background(), filter, update)
	return ret, err
}

func DeleteComment(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{"comment_id": id}

	ret, err := dao.MonCollection.DeleteOne(context.Background(), filter)
	return ret, err
}
