package controllers

import (
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
		logger.Error(map[string]interface{}{"[CreateComment] dao.MonCollection.InsertOne error:": err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": ret.InsertedID.(primitive.ObjectID)})
}
