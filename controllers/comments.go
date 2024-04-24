package controllers

import (
	"context"
	"net/http"
	mongodb "restApi/models/mongo"
	"restApi/pkg/logger"
	restapi "restApi/restApi"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentsController struct{}

func (com CommentsController) CreateComment(c *gin.Context) {
	var commentReq restapi.CommentReq
	if err := c.ShouldBindJSON(&commentReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment := mongodb.Comment{
		MediaID:     commentReq.MediaID,
		UserID:      commentReq.UserID,
		CommentText: commentReq.CommentText,
		CreatedAt:   time.Now(),
	}
	ret, err := mongodb.CreateComment(comment)

	if err != nil {
		logger.Error(map[string]interface{}{"[CreateComment] mongodb.CreateComment error:": err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": ret.InsertedID.(primitive.ObjectID)})
}

func (com CommentsController) GetCommentList(c *gin.Context) {
	cur, err := mongodb.GetCommentList()

	if err != nil {
		logger.Error(map[string]interface{}{"[GetCommentList] mongodb.GetCommentList error:": err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}
	defer cur.Close(context.Background())

	var comments []mongodb.Comment
	if err = cur.All(context.Background(), &comments); err != nil {
		logger.Error(map[string]interface{}{"[GetCommentList] cur.All error:": err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
