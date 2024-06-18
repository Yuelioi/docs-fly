package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Hitokoto struct {
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUID int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	CreatedAt  string `json:"created_at"`
	Length     int    `json:"length"`
}

func yiyan() (Hitokoto, error) {
	url := "https://v1.hitokoto.cn/?c=b"

	// 发送 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败:", err)
		return Hitokoto{}, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		return Hitokoto{}, err
	}

	var hitokoto Hitokoto
	json.Unmarshal(body, &hitokoto)

	// 输出响应体
	return hitokoto, nil
}

func GetYiYan(c *gin.Context) {
	hitokoto, _ := yiyan()
	sendSuccessResponse(c, time.Now(), hitokoto)
}
