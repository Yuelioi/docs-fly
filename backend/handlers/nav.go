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

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	var navs []models.NavData
	var cats []models.Category
	db.Model(&models.Category{}).Preload("Books").Find(&cats)

	// 遍历分类
	for _, cat := range cats {
		nav := models.NavData{
			MetaData: models.MetaData{
				Name:   cat.Name,
				Title:  cat.Title,
				Order:  cat.Order,
				Icon:   cat.Icon,
				Status: cat.Status,
			},
			Children: make([]models.MetaData, 0),
		}

		navs = append(navs, nav)
	}

	c.JSON(http.StatusOK, navs)

}
