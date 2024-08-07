package common

import (
	"docsfly/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var Responser = Response{}

type Response struct{}

func (Response) Success(c *gin.Context, clientTime time.Time, data interface{}) {
	c.JSON(http.StatusOK, models.ResponseBasicData{
		ClientTime: clientTime,
		IP:         c.ClientIP(),
		ServerTime: time.Now(),
		StatusCode: http.StatusOK,
		Data:       data,
	})
}

func (Response) Fail(c *gin.Context, statusCode int, clientTime time.Time, errMessage string) {
	c.JSON(statusCode, models.ResponseBasicData{
		ClientTime: clientTime,
		IP:         c.ClientIP(),
		ServerTime: time.Now(),
		StatusCode: statusCode,
		Data:       gin.H{"error": errMessage},
	})
}

func (Response) SuccessPageData(c *gin.Context, clientTime time.Time, data interface{}, totalCount int64, page, pageSize int) {
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

func (Response) FailPageData(c *gin.Context, statusCode int, clientTime time.Time, totalCount int64, page, pageSize int, errMessage string) {
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
