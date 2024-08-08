package comment

import (
	"docsfly/common"
	"docsfly/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetComments(c *gin.Context) {
	clientTime := time.Now()
	url := c.Query("url")
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var comments []models.Comment

	db.Model(models.Comment{}).Preload("Replies").Where("parent = 0").Scopes(common.MatchUrlPath(url)).Find(&comments)

	common.SendSuccessResponse(c, clientTime, comments)

}

func SendComment(c *gin.Context) {
	clientTime := time.Now()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, clientTime, "Invalid request payload")
		return
	}

	comment.IP = c.ClientIP()

	if comment.Nickname == "" || comment.Content == "" || comment.URL == "" {
		common.SendErrorResponse(c, http.StatusBadRequest, clientTime, "Invalid comment data")
		return
	}

	if err := db.Create(&comment).Error; err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, clientTime, "Failed to save comment,Database error")
		return
	}

	common.SendSuccessResponse(c, clientTime, gin.H{"message": "Comment posted successfully", "comment": comment})

}
