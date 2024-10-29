package controllers

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`           // 响应状态码
	Message string      `json:"message"`        // 消息或错误信息
	Data    interface{} `json:"data"`           // 数据
	Meta    MetaInfo    `json:"meta,omitempty"` // 元数据信息
}

type MetaInfo struct {
	IP string `json:"ip"`
}

func ReturnResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(200, Response{
		Code:    code,
		Message: message,
		Data:    data,
		Meta: MetaInfo{
			IP: c.ClientIP(),
		},
	})
}
func ReturnFailResponse(c *gin.Context, code int, message string) {
	c.JSON(200, Response{
		Code:    code,
		Message: message,
		Meta: MetaInfo{
			IP: c.ClientIP(),
		},
	})
}

func ReturnSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    0,
		Message: "",
		Data:    data,
		Meta: MetaInfo{
			IP: c.ClientIP(),
		},
	})
}
