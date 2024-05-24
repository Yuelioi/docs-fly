package routes

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

func DecodeQueryParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 遍历 URL 查询参数，对每个参数进行解码
		for key, value := range c.Request.URL.Query() {
			decodedValue, err := url.QueryUnescape(value[0])
			if err == nil {
				c.Request.URL.Query()[key] = []string{decodedValue}
			}
		}
		c.Next()
	}

}

func SetupRouterPlugins(engine *gin.Engine) {

	// 跨域处理
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"} // 允许的源
	// config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Content-Length"}
	// engine.Use(cors.New(config))

	// 参数解码
	engine.Use(DecodeQueryParams())
}
