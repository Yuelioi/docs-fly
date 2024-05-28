package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRndName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ip": c.ClientIP(), "nickname": RndName()})
}

func GetRndPoem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ip": c.ClientIP(), "content": RndPoem()})
}
