package handlers

import (
	"bytes"
	"docsfly/database"
	"docsfly/models"
	"docsfly/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	slug := c.Query("slug")
	document := c.Query("document")
	db, err := database.DbManager.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	documentInfo := getPostData(db, slug, document)

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

func getPostData(db *gorm.DB, slug, document string) *models.Entry {
	var documentInfo models.Entry
	db.Model(&models.Entry{}).Where("filepath = ?", slug+"/"+document).First(&documentInfo)
	return &documentInfo
}

// 保存文章 数据库+本地
func SavePost(c *gin.Context) {
	slug := c.Query("slug")
	document := c.Query("document")
	content := c.Query("content")

	ok, err := TokenVerifyMiddleware(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	db, err := database.DbManager.Connect()
	var documentInfo *models.Entry

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	documentInfo = getPostData(db, slug, document)

	// 写入本地文件
	err_write := os.WriteFile(document, []byte(content), 0644)
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
func buildFolderTree(folder *models.Chapter, categories []models.Entry, documents []models.Entry) {
	folder.Documents = make([]models.MetaData, 0)
	folder.Children = make([]models.Chapter, 0)
	// TODO 可以优化 添加后删除该文件夹

	// 添加文件到当前文件夹
	for _, doc := range documents {
		if strings.HasPrefix(doc.Filepath, folder.Filepath+"/") && doc.Depth == folder.Depth+1 {
			folder.Documents = append(folder.Documents, doc.MetaData)
		}
	}

	// 添加子文件夹到当前文件夹
	for _, cat := range categories {
		if strings.HasPrefix(cat.Filepath, folder.Filepath+"/") && cat.Depth == folder.Depth+1 {
			childFolder := models.Chapter{
				MetaData: cat.MetaData,
			}
			buildFolderTree(&childFolder, categories, documents)
			folder.Children = append(folder.Children, childFolder)
		}
	}
}

// 获取当前书籍章节信息,没有章节则直接获取文档信息
func GetChapter(c *gin.Context) {
	slug := c.Query("slug")

	parts := strings.Split(slug, "/")
	var book string

	// 检查是否有足够的部分
	if len(parts) >= 3 {
		book = strings.Join(parts[:3], "/")
	} else {
		return
	}

	db, err := database.DbManager.Connect()

	var categories []models.Entry
	var documents []models.Entry

	db.Where("filepath LIKE ?", book+"%").Find(&categories)
	db.Where("filepath LIKE ?", book+"%").Find(&documents)

	// 创建根文件夹
	var rootEntry models.Entry
	db.Where("filepath = ? AND depth = ?", book, 2).First(&rootEntry)

	rootFolder := models.Chapter{
		MetaData:  rootEntry.MetaData,
		Documents: make([]models.MetaData, 0),
		Children:  make([]models.Chapter, 0),
	}

	// 构建文件夹树
	buildFolderTree(&rootFolder, categories, documents)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	c.JSON(http.StatusOK, rootFolder)
}
