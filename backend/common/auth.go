package common

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var SigningKey = []byte("docs.yuelili.com")

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

	token, err := checkToken(tokenString)

	if err != nil {
		return nil, err
	}
	return token, nil

}

// 验证token字符串
func checkToken(tokenString string) (*jwt.Token, error) {
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
