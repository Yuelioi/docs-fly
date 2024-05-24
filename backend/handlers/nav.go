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
				Identity:    cat.Identity,
				DisplayName: cat.DisplayName,
				Order:       cat.Order,
				Icon:        cat.Icon,
				Hidden:      cat.Hidden,
			},
			Children: make([]models.MetaData, 0),
		}

		// 遍历书籍
		for _, book := range cat.Books {
			bk := models.MetaData{
				Identity:    book.Identity,
				DisplayName: book.DisplayName,
				Order:       book.Order,
				Icon:        book.Icon,
				Hidden:      book.Hidden,
			}
			nav.Children = append(nav.Children, bk)
		}
		navs = append(navs, nav)
	}

	c.JSON(http.StatusOK, navs)

}
