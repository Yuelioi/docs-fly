package models

import (
	"time"

	"gorm.io/gorm"
)

// CommentType 分为 article 和 book
type Comment struct {
	gorm.Model
	IP          string
	CreatedAt   time.Time `gorm:"default:0" json:"created_at"`
	Nickname    string    `gorm:"type:text;not null" json:"nickname"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	Parent      uint      `gorm:"default:0" json:"parent"`                       // 父评论 ID，顶级评论为 0
	Replies     []Comment `gorm:"foreignKey:Parent" json:"replies"`              // 通过外键关联回复
	CommentType string    `gorm:"type:varchar(20);not null" json:"comment_type"` // 评论类型，文章评论或书籍评论
}
