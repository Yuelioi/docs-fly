package handlers

import (
	"bytes"
	"docsfly/database"
	"docsfly/models"
	"docsfly/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"gorm.io/gorm"
)

func GetDepth(node atom.Atom) int {
	switch node {
	case atom.H1:
		return 1
	case atom.H2:
		return 2
	case atom.H3:
		return 3
	case atom.H4:
		return 3
	default:
		return 0
	}

}

func loopNode(n *html.Node, entries *[]models.Toc) {
	if n.Type == html.ElementNode && (n.DataAtom == atom.H1 || n.DataAtom == atom.H2) {
		for _, a := range n.Attr {
			if a.Key == "id" {
				*entries = append(*entries, models.Toc{
					ID:    a.Val,
					Depth: uint(GetDepth(n.DataAtom)),
					Title: textContent(n),
				})
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		loopNode(c, entries)
	}
}

func GenerateTOC(htmlContent string) ([]byte, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	var entries []models.Toc

	loopNode(doc, &entries)

	jsonData, err := json.Marshal(entries)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return jsonData, err
	}
	return jsonData, nil
}

func textContent(n *html.Node) string {
	var b bytes.Buffer
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			b.WriteString(c.Data)
		}
	}
	return b.String()
}

type PostResponseData struct {
	ContentMarkdown string `json:"content_markdown"`
	ContentHTML     string `json:"content_html"`
	TOC             string `json:"toc"`
}

// 文章页面获取文章markdown
func GetPost(c *gin.Context) {

	category := c.Query("category")
	book := c.Query("book")
	locale := c.Query("locale")
	chapter := c.Query("chapter")
	section := c.Query("section")
	document := c.Query("document")

	db, err := database.DbManager.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	documentInfo := getPostData(db, category, book, locale, chapter, section, document)

	htmlContent := string(utils.MarkdownToHTML([]byte(documentInfo.Content)))

	toc, _ := GenerateTOC(htmlContent)

	responseData := PostResponseData{
		ContentMarkdown: documentInfo.Content,
		ContentHTML:     htmlContent,
		TOC:             string(toc),
	}
	if htmlContent == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "No documents found"})
		return
	}

	c.JSON(http.StatusOK, responseData)

}

// 文章Markdown转Html
func GetPostHtml(c *gin.Context) {
	content := c.Query("content")

	htmlContent := string(utils.MarkdownToHTML([]byte(content)))

	toc, _ := GenerateTOC(htmlContent)
	responseData := PostResponseData{
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

func getPostData(db *gorm.DB, category, book, locale, chapter, section, document string) *models.Document {
	var catInfo models.Category

	var documentInfo models.Document

	db.Model(&catInfo).Select("id").Where("identity = ?", category).First(&catInfo)

	return &documentInfo

}

// 保存文章 数据库+本地
func SavePost(c *gin.Context) {
	category := c.Query("category")
	book := c.Query("book")
	locale := c.Query("locale")
	chapter := c.Query("chapter")
	section := c.Query("section")
	document := c.Query("document")

	content := c.Query("content")

	ok, err := TokenVerifyMiddleware(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	db, err := database.DbManager.Connect()
	var documentInfo *models.Document

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	documentInfo = getPostData(db, category, book, locale, chapter, section, document)

	// 写入本地文件
	err_write := utils.WriteMarkdownFile(category, book, locale, chapter, section, document, content)
	if err_write != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err_write.Error()})
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
	toc, err := GenerateTOC(htmlContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseData := PostResponseData{
		ContentMarkdown: documentInfo.Content,
		ContentHTML:     htmlContent,
		TOC:             string(toc),
	}

	c.JSON(http.StatusOK, &responseData)

}

// 获取当前书籍章节信息,没有章节则直接获取文档信息
func GetChapter(c *gin.Context) {
	category := c.Query("category")

	db, err := database.DbManager.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	var catInfo models.Category

	db.Model(&catInfo).Where("identity = ?", category).First(&catInfo)

	var result []models.ChapterInfo

	// 情况2.有章节

	// 情况2.1 章节目录层级也有文章, 即父级chapter为0的文章
	var chapterDocs []models.Document

	id := uint(1)
	for _, document := range chapterDocs {

		result = append(result, models.ChapterInfo{
			Category: catInfo.MetaData,
			Document: document.MetaData,
			ID:       id},
		)
		id += 1
	}

	// 情况2.2 获取章节

	c.JSON(http.StatusOK, result)
}
