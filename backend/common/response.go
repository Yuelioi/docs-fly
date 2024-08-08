package common

import (
	"docsfly/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SendSuccessResponse(c *gin.Context, clientTime time.Time, data interface{}) {
	c.JSON(http.StatusOK, models.ResponseBasicData{
		ClientTime: clientTime,
		IP:         c.ClientIP(),
		ServerTime: time.Now(),
		StatusCode: http.StatusOK,
		Data:       data,
	})
}

func SendErrorResponse(c *gin.Context, statusCode int, clientTime time.Time, errMessage string) {

	c.JSON(statusCode, models.ResponseBasicData{
		ClientTime: clientTime,
		IP:         c.ClientIP(),
		ServerTime: time.Now(),
		StatusCode: statusCode,
		Data:       gin.H{"error": errMessage},
	})
	c.Abort()
}

func SendSuccessResponsePageData(c *gin.Context, clientTime time.Time, data interface{}, totalCount int64, page, pageSize int) {
	c.JSON(http.StatusOK, models.ResponsePageData{
		ClientTime: clientTime,
		IP:         c.ClientIP(),
		ServerTime: time.Now(),
		StatusCode: http.StatusOK,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		Data:       data,
	})
}

func SendErrorResponsePageData(c *gin.Context, statusCode int, clientTime time.Time, totalCount int64, page, pageSize int, errMessage string) {
	c.JSON(statusCode, models.ResponsePageData{
		ClientTime: clientTime,
		IP:         c.ClientIP(),
		ServerTime: time.Now(),
		StatusCode: statusCode,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		Data:       gin.H{"error": errMessage},
	})
}
