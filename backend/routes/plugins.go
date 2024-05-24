package routes

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
//
// 注意：如果一个查询参数有多个值（例如，key=value1&key=value2）
// 只解码第一个值（即 value1）。
// 其他值（例如 value2）将保持原样。
func DecodeQueryParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		for key, value := range c.Request.URL.Query() {
			decodedValue, err := url.QueryUnescape(value[0])
			if err == nil {
				c.Request.URL.Query()[key] = []string{decodedValue}
			}
		}
		c.Next()
	}
}
