package handlers

import (
	"docsfly/database"
	"docsfly/models"
	"docsfly/utils"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var SigningKey = []byte("docs.yuelili.com")

// 登录验证,如果成功则返回token
func LoginAuth(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	var databasePassword string

	db, err := database.DbManager.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	db.Model(models.User{}).Where("username = ?", username).Pluck("password", &databasePassword)

	if utils.CheckPasswordHash(password, databasePassword) {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["authorized"] = true
		claims["username"] = username
		claims["role"] = "admin"
		claims["exp"] = time.Now().Add(time.Hour).Unix()

		tokenString, err := token.SignedString(SigningKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token解析失败"})
			return
		}
		c.JSON(http.StatusOK, tokenString)

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "账号密码错误"})
	}

}

// 解析token
func parseToken(c *gin.Context) (*jwt.Token, error) {
	// Extract token from the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("missing authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	tokenString := parts[1]

	token, err := CheckToken(tokenString)

	if err != nil {
		return nil, err
	}
	return token, nil

}

// 验证token字符串
func CheckToken(tokenString string) (*jwt.Token, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return SigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func TokenVerify(c *gin.Context) {
	token := c.Query("token")

	_, err := CheckToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})

}

// 验证token中间件
func TokenVerifyMiddleware(c *gin.Context) (bool, error) {
	// 解析Token
	token, err := parseToken(c)
	if err != nil {
		c.Abort()
		return false, err
	}

	// 验证是否有效
	if !token.Valid {
		c.Abort()
		return false, errors.New("invalid token")

	}

	return true, nil

}
