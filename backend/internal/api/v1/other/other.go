package other

import (
	"docsfly/internal/common"
	"docsfly/internal/global"
	"docsfly/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func VisitorInsertLog(c *gin.Context) {

	clientTime := time.Now()
	url := c.Query("url")
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}
	db := dbContext.(*gorm.DB)

	var count int64
	db.Scopes(common.BasicModel, common.MatchUrlPath(url)).Count(&count)

	if count == 0 {
		common.Responser.Fail(c, http.StatusInternalServerError, clientTime, "Can't find target link")
		return
	}

	urlList := strings.Split(url, "/")

	var category, book, locale string

	if len(urlList) > 2 {
		category = urlList[0]
		book = urlList[1]
		locale = urlList[2]
	} else {
		common.Responser.Fail(c, http.StatusInternalServerError, clientTime, "Can't find target path")
		return
	}

	today := time.Now().Local()

	vs := models.Visitor{
		IP:       c.ClientIP(),
		URL:      url,
		Time:     today,
		Category: category,
		Book:     book,
		Locale:   locale,
	}

	db.Model(&models.Visitor{}).Create(&vs)

	// 返回 IP 地址给客户端
	common.Responser.Success(c, clientTime, gin.H{"message": "success"})
}

func GetAppVersion(c *gin.Context) {
	common.Responser.Success(c, time.Now(), global.AppConfig.AppConfig.AppVersion)
}

func GetRndName(c *gin.Context) {
	common.Responser.Success(c, time.Now(), rndName())
}

func GetRndPost(c *gin.Context) {
	clientTime := time.Now()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}
	var doc models.Entry

	db := dbContext.(*gorm.DB)
	if err := db.Scopes(common.BasicModel, common.FindFile).Order("RANDOM()").First(&doc).Error; err != nil {
		common.Responser.Fail(c, http.StatusInternalServerError, clientTime, "Could not retrieve a random post")
		return
	}

	common.Responser.Success(c, time.Now(), doc.MetaData)
}

func GetRndPoem(c *gin.Context) {
	common.Responser.Success(c, time.Now(), rndPoem())
}

type Hitokoto struct {
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUID int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	CreatedAt  string `json:"created_at"`
	Length     int    `json:"length"`
}

func yiyan() (Hitokoto, error) {
	url := "https://v1.hitokoto.cn/?c=b"

	// 发送 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败:", err)
		return Hitokoto{}, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		return Hitokoto{}, err
	}

	var hitokoto Hitokoto
	json.Unmarshal(body, &hitokoto)

	// 输出响应体
	return hitokoto, nil
}

func GetYiYan(c *gin.Context) {
	hitokoto, _ := yiyan()
	common.Responser.Success(c, time.Now(), hitokoto)
}
