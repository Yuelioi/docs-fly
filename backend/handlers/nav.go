package handlers

import (
	"docsfly/database"
	"docsfly/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取顶部导航栏信息
func GetNav(c *gin.Context) {

	db, err := database.DbManager.Connect()

	navs := models.MetaData{}
	db.Model(models.MetaData{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	c.JSON(http.StatusOK, navs)

}
