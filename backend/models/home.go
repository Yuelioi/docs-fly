package models

// 主页统计信息
type Statistic struct {
	BookCount              int64 `json:"book_count"`
	DocumentCount          int64 `json:"document_count"`
	HistoricalVisitorCount int64 `json:"historical_visitor_count"`
	TodayVisitorCount      int64 `json:"today_visitor_count"`
}

type SearchOption struct {
	MetaData
	Children []MetaData `json:"children"`
}

type SearchResult struct {
	SearchTime string       `json:"search_time"`
	Result     []SearchData `json:"result"`
}

// 主页搜索显示数据
type SearchData struct {
	Locale string `json:"locale"`

	CategoryIdentity    string `json:"category_identity"`
	CategoryDisplayName string `json:"category_display_name"`

	BookIdentity    string `json:"book_identity"`
	BookDisplayName string `json:"book_display_name"`

	ChapterIdentity    string `json:"chapter_identity"`
	ChapterDisplayName string `json:"chapter_display_name"`

	SectionIdentity    string `json:"section_identity"`
	SectionDisplayName string `json:"section_display_name"`

	DocumentIdentity    string `json:"document_identity"`
	DocumentDisplayName string `json:"document_display_name"`

	Content string `json:"content"`
}
