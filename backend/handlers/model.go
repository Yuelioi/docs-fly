package handlers

import (
	"docsfly/models"
	"time"
	// 假设这是你的数据库包
)

type BookData struct {
	Url         string          `json:"url"`
	ChapterType string          `json:"chapter_type"`
	MetaData    models.MetaData `json:"metadata"`
}

// 主页统计信息
type Statistic struct {
	BookCount              int64 `json:"book_count"`
	DocumentCount          int64 `json:"document_count"`
	HistoricalVisitorCount int64 `json:"historical_visitor_count"`
	TodayVisitorCount      int64 `json:"today_visitor_count"`
}

type SearchResult struct {
	SearchTime string       `json:"search_time"`
	Result     []SearchData `json:"result"`
}

// 主页搜索显示数据
type SearchData struct {
	Url           string `json:"url"`
	Locale        string `json:"locale"`
	CategoryTitle string `json:"category_title"`
	BookTitle     string `json:"book_title"`
	DocumentTitle string `json:"document_title"`
	Content       string `json:"content"`
}

type Nav struct {
	MetaData models.MetaData   `json:"metadata"`
	Children []models.MetaData `json:"children"`
}

// ResponseData 结构体用于扩展返回数据
type ResponseData struct {
	MetaData   []Nav     `json:"metadata"`
	Timestamp  time.Time `json:"timestamp"`
	IP         string    `json:"ip"`
	ServerTime time.Time `json:"server_time"`
	StatusCode int       `json:"status_code"`
}

// 文章目录
type Toc struct {
	ID    string `json:"id"`
	Depth uint   `json:"depth"`
	Title string `json:"title"`
}

// 一个章节的信息
type Chapter struct {
	MetaData  models.MetaData `json:"metadata"`
	Filepath  string
	Documents []models.MetaData `json:"documents"`
	Children  []Chapter         `json:"children"`
}
