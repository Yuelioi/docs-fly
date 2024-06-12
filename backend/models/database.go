// 数据库模型
package models

import (
	"time"

	"gorm.io/gorm"
)

// 文件或者文件夹信息
type Entry struct {
	gorm.Model
	MetaData
	ModTime time.Time
	IsDir   bool
	Locale  string
	Content string
}

// 访客记录
type Visitor struct {
	gorm.Model
	IP       string    // 访客的 IP 地址
	URL      string    // 访客访问的 URL
	Time     time.Time // 访问时间
	Category string    // 分类 identity
	Book     string    // 书籍 identity
	Locale   string    // 语言
}

// 用户
type User struct {
	gorm.Model
	Username string `gorm:"unique_index"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	IP       string `json:"ip"`
}
