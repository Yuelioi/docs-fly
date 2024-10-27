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
	Locale  string
	Content string
}

// 访客记录
type Visitor struct {
	gorm.Model
	IP       string    // 访客的 IP 地址
	URL      string    `json:"url"` // 访客访问的 URL
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

// 评论系统, 不会有多级嵌套
type Comment struct {
	gorm.Model `json:"-"`
	ID         uint      `json:"id"`
	IP         string    `json:"-"`
	CreatedAt  time.Time ` json:"createdAt"`
	Nickname   string    `gorm:"type:text;not null" json:"nickname"`
	Content    string    `gorm:"type:text;not null" json:"content"`
	Parent     uint      `gorm:"default:0" json:"parent"`          // 父评论 ID，默认为0 无父级
	Replies    []Comment `gorm:"foreignKey:Parent" json:"replies"` // 通过外键关联回复
	URL        string    `gorm:"type:text;not null" json:"url"`    // 评论所属的链接地址
}

func (c *Comment) AfterFind(tx *gorm.DB) (err error) {
	if c.Replies == nil {
		c.Replies = []Comment{}
	}
	return
}
