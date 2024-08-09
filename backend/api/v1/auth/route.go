package auth

import (
	"docsfly/internal/common"
	"docsfly/internal/global"
	"docsfly/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthRouter struct{}

func (ar *AuthRouter) Register(engine *gin.Engine) {
	engine.POST("/"+global.AppConfig.APIVersion+"/auth/login", ar.LoginAuth)

}

// 登录验证,如果成功则返回token
func (ar *AuthRouter) LoginAuth(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	var databasePassword string

	clientTime := time.Now()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	db.Model(models.User{}).Where("username = ?", username).Pluck("password", &databasePassword)

	if checkPasswordHash(password, databasePassword) {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["authorized"] = true
		claims["username"] = username
		claims["role"] = "admin"
		claims["exp"] = time.Now().Add(time.Hour).Unix()

		tokenString, err := token.SignedString(common.SigningKey)
		if err != nil {
			common.Responser.Fail(c, http.StatusUnauthorized, clientTime, "token解析失败")
			return
		}
		common.Responser.Success(c, clientTime, tokenString)

	} else {
		common.Responser.Fail(c, http.StatusUnauthorized, clientTime, "账号密码错误")
		return
	}
}
