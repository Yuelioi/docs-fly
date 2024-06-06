// 数据库模型
package models

import (
	"time"

	"gorm.io/gorm"
)

// 用于写入本地
type MetaDataLocal struct {
	Size     int
	Hash     string
	Filepath string
}

// 分类, 记录文件夹
type Category struct {
	gorm.Model
	MetaData
	ModTime   time.Time
	Display   bool // 是否显示
	Documents []Document
}

// 文档(.md) 记录文件信息
type Document struct {
	gorm.Model
	MetaData
	ModTime    time.Time
	Locale     string
	Content    string
	CategoryID uint
	Category   Category
	Status     string
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
