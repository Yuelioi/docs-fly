package router

import (
	"docsfly/internal/database"
	"docsfly/models"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

// 注册插件
func registerPlugins(engine *gin.Engine) {
	// 参数解码
	engine.Use(DecodeQueryParams())
	engine.Use(DBMiddleware())
}

// 解码 URL 编码的查询参数。
func DecodeQueryParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Request.URL.RawQuery
		decodedQuery, err := url.QueryUnescape(query)
		if err == nil {
			c.Request.URL.RawQuery = decodedQuery
		}
		c.Next()
	}
}

// 连接数据库
func DBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.DbManager.Connect()

		if err != nil {
			c.JSON(http.StatusOK, models.ResponseBasicData{
				ClientTime: time.Now(),
				IP:         c.ClientIP(),
				ServerTime: time.Now(),
				StatusCode: http.StatusBadRequest,
				Data:       gin.H{"error": "Cannot Connect Database"},
			})
			c.Abort()
			return
		}

		c.Set("db", db)

		c.Next()
	}
}
