package handlers

import (
	"bytes"
	"docsfly/global"
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

func getDepth(node atom.Atom) int {
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

func loopNode(n *html.Node, entries *[]Toc) {
	if n.Type == html.ElementNode && (n.DataAtom == atom.H1 || n.DataAtom == atom.H2) {
		for _, a := range n.Attr {
			if a.Key == "id" {
				*entries = append(*entries, Toc{
					ID:    a.Val,
					Depth: uint(getDepth(n.DataAtom)),
					Title: textContent(n),
				})
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		loopNode(c, entries)
	}
}

func generateTOC(htmlContent string) ([]byte, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	var entries []Toc

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

// 文章页面获取文章markdown
func GetPost(c *gin.Context) {
	postPath := c.Query("postPath")
	clientTime := currentTime()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var documentInfo models.Entry
	db.Scopes(BasicModel, MatchPath(postPath)).First(&documentInfo)

	htmlContent := string(utils.MarkdownToHTML([]byte(documentInfo.Content)))

	if htmlContent == "" {
		sendErrorResponse(c, http.StatusNotFound, clientTime, "No documents found")

		return
	}

	toc, _ := generateTOC(htmlContent)

	responseData := PostResponseBasicData{
		ContentMarkdown: documentInfo.Content,
		ContentHTML:     htmlContent,
		TOC:             string(toc),
	}

	sendSuccessResponse(c, clientTime, responseData)

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

	ok, err := TokenVerifyMiddleware(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	clientTime := currentTime()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var documentInfo models.Entry
	db.Scopes(BasicModel, MatchPath(postPath)).First(&documentInfo)

	documentPath := global.AppConfig.Resource + "/" + documentInfo.Filepath

	// 写入本地文件
	err_write := os.WriteFile(documentPath, []byte(content), 0644)

	if err_write != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, err_write.Error())

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
	sendSuccessResponse(c, clientTime, responseData)

}
func buildFolderTree(folder *Chapter, categories []models.Entry, documents []models.Entry) {
	folder.Documents = make([]models.MetaData, 0)
	folder.Children = make([]Chapter, 0)
	// TODO 可以优化 添加后删除该文件夹

	// 添加文件到当前文件夹
	for _, doc := range documents {
		if strings.HasPrefix(doc.Filepath, folder.Filepath+"/") && doc.Depth == folder.MetaData.Depth+1 {
			folder.Documents = append(folder.Documents, doc.MetaData)
		}
	}

	// 添加子文件夹到当前文件夹
	for _, cat := range categories {
		if strings.HasPrefix(cat.Filepath, folder.Filepath+"/") && cat.Depth == folder.MetaData.Depth+1 {
			childFolder := Chapter{
				MetaData: cat.MetaData,
			}
			buildFolderTree(&childFolder, categories, documents)
			folder.Children = append(folder.Children, childFolder)
		}
	}
}

// 获取当前书籍章节信息,没有章节则直接获取文档信息
func GetChapter(c *gin.Context) {
	postPath := c.Query("postPath")

	clientTime := currentTime()
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
		return
	}

	var categories []models.Entry
	var documents []models.Entry

	db.Scopes(BasicModel, HasPrefixPath(book)).Find(&categories)
	db.Scopes(BasicModel, HasPrefixPath(book)).Find(&documents)

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
	buildFolderTree(&chapterMeta, categories, documents)

	sendSuccessResponse(c, clientTime, chapterMeta)
}
