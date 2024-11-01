package controllers

import (
	"docsfly/internal/common"
	"docsfly/internal/config"
	"docsfly/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
}

func (cr *CommentController) Register(engine *gin.Engine) {
	engine.GET("/"+config.Instance.App.ApiVersion+"/comment", GetComments)
	engine.POST("/"+config.Instance.App.ApiVersion+"/comment", SendComment)
}

func GetComments(c *gin.Context) {
	url := c.Query("url")
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var comments []models.Comment

	db.Model(models.Comment{}).Preload("Replies").Where("parent = 0").Scopes(common.MatchUrlPath(url)).Find(&comments)

	ReturnSuccessResponse(c, comments)
}

func SendComment(c *gin.Context) {
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		ReturnFailResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	comment.IP = c.ClientIP()

	if comment.Nickname == "" || comment.Content == "" || comment.URL == "" {
		ReturnFailResponse(c, http.StatusBadRequest, "Invalid comment data")
		return
	}

	if err := db.Create(&comment).Error; err != nil {
		ReturnFailResponse(c, http.StatusBadRequest, "Failed to save comment,Database error")
		return
	}

	ReturnSuccessResponse(c, comment)

}
