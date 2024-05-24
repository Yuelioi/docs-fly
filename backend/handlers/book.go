package handlers

// BookPage

import (
	"docsfly/database"
	"docsfly/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取当前书籍章节信息,没有章节则直接获取文档信息
func GetBook(c *gin.Context) {
	category := c.Query("category")
	book := c.Query("book")
	locale := c.Query("locale")

	db, err := database.DbManager.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	var catInfo models.Category
	var bookInfo models.Book
	var chapterInfo models.Chapter

	db.Model(&catInfo).Where("identity = ?", category).First(&catInfo)
	db.Model(&bookInfo).Where("identity = ?", book).Where("category_id = ?", catInfo.ID).First(&bookInfo)
	db.Model(&models.Chapter{}).Where("book_id = ?", bookInfo.ID).First(&chapterInfo)

	result := []models.BookChapter{}

	if chapterInfo.ID == 0 {
		//	没有章节
		var documents []models.Document
		db.Model(&models.Document{}).Where("book_id = ?", bookInfo.ID).Find(&documents)

		if len(documents) > 0 {
			for _, document := range documents {
				result = append(result, models.BookChapter{
					MetaData: document.MetaData,
					Locale:   locale,
					Document: document.Identity,
				},
				)
			}
		}

	} else { // 有章节
		var chapters []models.Chapter

		// 查询章节并只选择指定的列
		db.Preload("Documents").Preload("Sections").Where("book_id = ?", bookInfo.ID).Where("locale = ?", locale).Where("locale=?", locale).Find(&chapters)

		if len(chapters) > 0 {
			// 循环遍历每个章节, 是否要获取文章信息 待定
			for _, chapter := range chapters {

				firstDocs := chapter.Documents[0]

				result = append(result, models.BookChapter{
					MetaData: chapter.MetaData,
					Chapter:  chapter.Identity,
					Section:  firstDocs.Section.Identity,
					Document: firstDocs.Identity,
				})
			}
		}

	}

	output := models.BookData{
		Category: catInfo.MetaData,
		Book:     bookInfo.MetaData,
		Children: result,
	}

	c.JSON(http.StatusOK, output)
}
