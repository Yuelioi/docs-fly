package models

type BookChapter struct {
	MetaData
	Locale   string `json:"locale"`
	Chapter  string `json:"chapter"`
	Section  string `json:"section"`
	Document string `json:"document"`
}

// 书籍页面数据
type BookData struct {
	Category MetaData      `json:"category"`
	Book     MetaData      `json:"book"`
	Children []BookChapter `json:"children"`
}
