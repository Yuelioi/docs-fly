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

	CategoryName  string `json:"category_identity"`
	CategoryTitle string `json:"category_display_name"`

	BookName  string `json:"book_identity"`
	BookTitle string `json:"book_display_name"`

	ChapterName  string `json:"chapter_identity"`
	ChapterTitle string `json:"chapter_display_name"`

	SectionName  string `json:"section_identity"`
	SectionTitle string `json:"section_display_name"`

	DocumentName  string `json:"document_identity"`
	DocumentTitle string `json:"document_display_name"`

	Content string `json:"content"`
}
