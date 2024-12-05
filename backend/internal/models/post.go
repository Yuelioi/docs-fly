package models

type PostResponseBasicData struct {
	ContentMarkdown string `json:"content_markdown"`
	ContentHTML     string `json:"content_html"`
	TOC             string `json:"toc"`
}

// 一个章节的信息
type Chapter struct {
	MetaData  MetaData   `json:"metadata"`
	Filepath  string     `json:"filepath"`
	Documents []MetaData `json:"documents"`
	Children  []Chapter  `json:"children"`
}

// 文章目录
type Toc struct {
	ID    string `json:"id"`
	Depth uint   `json:"depth"`
	Title string `json:"title"`
}
