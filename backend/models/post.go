package models

// 1.只有文章  Chapter.id == 0

// 2.有章节 有文章

// 3. 有章节 有小节 有文章

// 文章页左侧大纲
type ChapterInfo struct {
	Category  MetaData   `json:"category"`
	Chapter   MetaData   `json:"chapter"`
	Sections  []MetaData `json:"sections"`
	Document  MetaData   `json:"document"`
	Documents []MetaData `json:"documents"`

	ID uint `json:"id"`
}

// 文章目录
type Toc struct {
	ID    string `json:"id"`
	Depth uint   `json:"depth"`
	Title string `json:"title"`
}
