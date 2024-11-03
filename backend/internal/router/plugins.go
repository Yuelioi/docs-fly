package router

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

// 注册插件
func registerPlugins(engine *gin.Engine) {
	// 参数解码
	engine.Use(DecodeQueryParams())
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
