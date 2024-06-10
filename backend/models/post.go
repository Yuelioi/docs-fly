package models

// 文章目录
type Toc struct {
	ID    string `json:"id"`
	Depth uint   `json:"depth"`
	Title string `json:"title"`
}

// 一个章节的信息
type Chapter struct {
	MetaData  `json:"metadata"`
	Documents []MetaData `json:"documents"`
	Children  []Chapter  `json:"children"`
}
