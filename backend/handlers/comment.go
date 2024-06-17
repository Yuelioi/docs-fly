package handlers

import (
	"docsfly/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetComments(c *gin.Context) {
	clientTime := currentTime()
	url := c.Query("url")
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var comments []models.Comment

	db.Model(models.Comment{}).Preload("Replies").Where("parent = 0").Scopes(MatchUrlPath(url)).Find(&comments)

	sendSuccessResponse(c, clientTime, comments)

}

func SendComment(c *gin.Context) {
	clientTime := currentTime()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, clientTime, "Invalid request payload")
		return
	}

	comment.IP = c.ClientIP()

	if comment.Nickname == "" || comment.Content == "" || comment.URL == "" {
		sendErrorResponse(c, http.StatusBadRequest, clientTime, "Invalid comment data")
		return
	}

	if err := db.Create(&comment).Error; err != nil {
		sendErrorResponse(c, http.StatusBadRequest, clientTime, "Failed to save comment,Database error")
		return
	}

	sendSuccessResponse(c, clientTime, gin.H{"message": "Comment posted successfully", "comment": comment})

}
