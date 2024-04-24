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
	ret, err := mongodb.CreateComment(&comment)

	if err != nil {
		logger.Error(map[string]interface{}{"[CreateComment] mongodb.CreateComment error:": err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": ret.InsertedID.(primitive.ObjectID), "comment": comment})
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

func (com CommentsController) UpdateComment(c *gin.Context) {
	var updatedComment mongodb.Comment
	if err := c.ShouldBindJSON(&updatedComment); err != nil {
		logger.Error(map[string]interface{}{"[UpdateComment] c.ShouldBindJSON error:": err.Error()})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedComment.CommentID = c.Param("id")
	updatedComment.CreatedAt = time.Now()

	_, err := mongodb.UpdateComment(updatedComment)

	if err != nil {
		logger.Error(map[string]interface{}{"[UpdateComment] mongodb.UpdateComment error:": err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

func (com CommentsController) DeleteComment(c *gin.Context) {
	commentId := c.Param("id")

	_, err := mongodb.DeleteComment(commentId)
	if err != nil {
		logger.Error(map[string]interface{}{"[DeleteComment] mongodb.DeleteComment error:": err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
