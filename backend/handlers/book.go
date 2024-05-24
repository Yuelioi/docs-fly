package handlers

// BookPage

import (
	"docsfly/database"
	"docsfly/models"
	"docsfly/utils"
	"encoding/json"
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

func GetBookMeta(c *gin.Context) {
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

	db.Model(&catInfo).Where("identity = ?", category).First(&catInfo)
	db.Model(&bookInfo).Where("identity = ?", book).Where("category_id = ?", catInfo.ID).First(&bookInfo)

	filepath := bookInfo.Filepath + "/" + locale + "/" + "meta.json"
	data, err := utils.ReadJson(filepath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load data"})
		return
	}

	c.JSON(http.StatusOK, &data)
}

func SaveBookMeta(c *gin.Context) {
	category := c.Query("category")
	book := c.Query("book")
	locale := c.Query("locale")
	metas_data := c.Query("metas")

	// 解析meta数据
	var metas []models.MetaData

	if err := json.Unmarshal([]byte(metas_data), &metas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	db, err := database.DbManager.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Parse Data"})
		return
	}

	var catInfo models.Category
	var bookInfo models.Book

	db.Model(&catInfo).Where("identity = ?", category).First(&catInfo)
	db.Model(&bookInfo).Preload("Chapters").Where("identity = ?", book).Where("category_id = ?", catInfo.ID).First(&bookInfo)

	// 保存meta.json
	filepath := bookInfo.Filepath + "/" + locale + "/" + "meta.json"
	err = utils.WriteJson(filepath, metas_data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Save Data"})
		return
	}

	// 更新数据库

	// 使用一个map来快速查找meta数据
	metaMap := make(map[string]models.MetaData)
	for _, meta := range metas {
		metaMap[meta.Identity] = meta
	}
	chaptersToUpdate := make([]models.Chapter, 0)

	for _, chapter := range bookInfo.Chapters {
		if meta, exists := metaMap[chapter.Identity]; exists {
			needsUpdate := false
			if chapter.DisplayName != meta.DisplayName {
				chapter.DisplayName = meta.DisplayName
				needsUpdate = true
			}
			if chapter.Hidden != meta.Hidden {
				chapter.Hidden = meta.Hidden
				needsUpdate = true
			}
			if chapter.Order != meta.Order {
				chapter.Order = meta.Order
				needsUpdate = true
			}
			if needsUpdate {
				chaptersToUpdate = append(chaptersToUpdate, chapter)
			}
		}
	}

	// 批量更新有变化的章节
	if len(chaptersToUpdate) > 0 {
		db.Save(&chaptersToUpdate)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data saved successfully"})
}
