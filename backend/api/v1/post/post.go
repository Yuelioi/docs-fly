package post

import (
	"docsfly/internal/common"
	"docsfly/internal/global"
	"docsfly/models"
	"docsfly/pkg/utils"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 文章页面获取文章markdown
func GetPost(c *gin.Context) {
	postPath := c.Query("postPath")
	clientTime := time.Now()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	// TODO 没有文章 就尝试切换语言

	var entryInfo models.Entry
	var htmlContent string
	db.Scopes(common.BasicModel, common.MatchUrlPath(postPath)).First(&entryInfo)

	if entryInfo.IsDir && entryInfo.Content == "" {

		var cats []models.Entry
		db.Scopes(common.BasicModel, common.HasPrefixUrlPath(postPath), common.FindFolder).Where("depth = ?", entryInfo.Depth+1).Find(&cats)

		var docs []models.Entry
		db.Scopes(common.BasicModel, common.HasPrefixUrlPath(postPath), common.FindFile).Where("depth = ?", entryInfo.Depth+1).Find(&docs)

		htmlContent = "<h2>小节</h2><ul>"

		for _, cat := range cats {
			htmlContent += fmt.Sprintf("<li><a href=\"#/post/%s\">%s</a></li>", cat.URL, cat.Title)
		}

		htmlContent += "</ul><h2>文章</h2><ul>"

		// 生成文件部分的 HTML
		for _, doc := range docs {
			htmlContent += fmt.Sprintf("<li><a href=\"#/post/%s\">%s</a></li>", doc.URL, doc.Title)
		}

		htmlContent += "</ul></body></html>"

	} else {
		htmlContent = string(utils.MarkdownToHTML([]byte(entryInfo.Content)))
	}

	// 文件没有内容 文件夹没有子级 那就报错吧
	if htmlContent == "" {
		common.Responser.Fail(c, http.StatusNotFound, clientTime, "No documents found")
		return
	}

	toc, _ := generateTOC(htmlContent)

	responseData := PostResponseBasicData{
		ContentMarkdown: entryInfo.Content,
		ContentHTML:     htmlContent,
		TOC:             string(toc),
	}

	common.Responser.Success(c, clientTime, responseData)

}

// 文章Markdown转Html
func GetPostHtml(c *gin.Context) {
	content := c.Query("content")

	htmlContent := string(utils.MarkdownToHTML([]byte(content)))

	toc, _ := generateTOC(htmlContent)
	responseData := PostResponseBasicData{
		ContentMarkdown: "",
		ContentHTML:     htmlContent,
		TOC:             string(toc),
	}
	if htmlContent == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "No documents found"})
		return
	}

	c.JSON(http.StatusOK, responseData)

}

// 保存文章 数据库+本地
func SavePost(c *gin.Context) {
	postPath := c.Query("postPath")
	content := c.Query("content")
	clientTime := time.Now()
	ok, err := common.TokenVerifyMiddleware(c)

	if !ok {
		common.Responser.Fail(c, http.StatusUnauthorized, clientTime, err.Error())
		return
	}

	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var documentInfo models.Entry
	db.Scopes(common.BasicModel, common.MatchUrlPath(postPath)).First(&documentInfo)

	var documentPath string

	if documentInfo.IsDir {
		documentPath = global.AppConfig.Resource + "/" + documentInfo.Filepath + "/" + "README.md"
	} else {
		documentPath = global.AppConfig.Resource + "/" + documentInfo.Filepath

	}

	// 写入本地文件
	err_write := os.WriteFile(documentPath, []byte(content), 0644)

	if err_write != nil {
		common.Responser.Fail(c, http.StatusInternalServerError, clientTime, err_write.Error())

		return
	}

	// 写入数据库
	documentInfo.Content = content
	if err := db.Save(documentInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 重新获取html与toc
	htmlContent := string(utils.MarkdownToHTML([]byte(content)))
	toc, err := generateTOC(htmlContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseData := PostResponseBasicData{
		ContentMarkdown: documentInfo.Content,
		ContentHTML:     htmlContent,
		TOC:             string(toc),
	}
	common.Responser.Success(c, clientTime, responseData)

}

// 获取当前书籍章节信息,没有章节则直接获取文档信息
func GetChapter(c *gin.Context) {
	postPath := c.Query("postPath")

	clientTime := time.Now()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	parts := strings.Split(postPath, "/")
	var book string

	// 检查是否有足够的部分
	if len(parts) >= 3 {
		book = strings.Join(parts[:3], "/")
	} else {
		common.Responser.Fail(c, http.StatusBadRequest, clientTime, "Invalid post path")
	}

	// ok, cachedData := getCache(book)
	// if ok {
	// 	common.Responser.Success(c, clientTime, cachedData)
	// 	return
	// }

	var categories, documents, allEntries []models.Entry

	db.Scopes(common.BasicModel, common.HasPrefixUrlPath(book)).Find(&allEntries)

	// 需要忽略README文件
	for _, entry := range allEntries {
		if strings.ToLower(entry.Name) == "readme.md" {
			continue
		}
		if entry.IsDir {
			categories = append(categories, entry)
		} else {
			documents = append(documents, entry)
		}
	}

	// 创建根文件夹
	var rootEntry models.Entry
	db.Where("filepath = ? AND depth = ?", book, 2).First(&rootEntry)

	chapterMeta := Chapter{
		MetaData:  rootEntry.MetaData,
		Filepath:  rootEntry.MetaData.Filepath,
		Documents: make([]models.MetaData, 0),
		Children:  make([]Chapter, 0),
	}

	// 构建文件夹树
	buildFolderTree(&chapterMeta, categories, documents, 3)

	// 将数据存储到缓存中

	common.Responser.Success(c, clientTime, chapterMeta)
}
