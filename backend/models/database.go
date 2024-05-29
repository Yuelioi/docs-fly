package models

import (
	"time"

	"gorm.io/gorm"
)

// 用于写入本地 追加个临时的
type MetaDataLocal struct {
	MetaDatas []MetaData
	Filepath  string
}

type Category struct {
	gorm.Model
	MetaData
	ModTime  time.Time
	Books    []Book
	Filepath string
}

type Book struct {
	gorm.Model
	MetaData
	ModTime     time.Time
	Description string `json:"description"` // 从Readme读取
	CategoryID  uint
	Category    Category
	Chapters    []Chapter
	Sections    []Section
	Documents   []Document
	Filepath    string
}

type Chapter struct {
	gorm.Model
	MetaData
	ModTime    time.Time
	Locale     string
	CategoryID uint
	Category   Category
	BookID     uint
	Book       Book
	Sections   []Section
	Documents  []Document
	Filepath   string
}

type Section struct {
	gorm.Model
	MetaData
	ModTime    time.Time
	Locale     string
	CategoryID uint
	Category   Category
	BookID     uint
	Book       Book
	ChapterID  uint
	Chapter    Chapter
	Documents  []Document
	Filepath   string
}

type Document struct {
	gorm.Model
	MetaData
	ModTime    time.Time
	Locale     string
	Content    string
	Html       string
	CategoryID uint
	Category   Category
	BookID     uint
	Book       Book
	ChapterID  uint
	Chapter    Chapter
	Section    Section
	SectionID  uint
	Filepath   string
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
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	IP       string `json:"ip"`
}
