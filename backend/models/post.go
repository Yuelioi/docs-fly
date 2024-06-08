package models

// 文章目录
type Toc struct {
	ID    string `json:"id"`
	Depth uint   `json:"depth"`
	Title string `json:"title"`
}
