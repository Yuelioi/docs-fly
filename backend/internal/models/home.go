package models

type SearchResult struct {
	Result []SearchData `json:"result"`
}

// 主页搜索显示数据
type SearchData struct {
	Index         int    `json:"index"` // 从1开始!
	Url           string `json:"url"`
	Locale        string `json:"locale"`
	CategoryTitle string `json:"category_title"`
	BookTitle     string `json:"book_title"`
	DocumentTitle string `json:"document_title"`
	Content       string `json:"content"`
}

// 主页统计信息
type HomeStatistic struct {
	BookCount              int64 `json:"book_count"`
	DocumentCount          int64 `json:"document_count"`
	HistoricalVisitorCount int64 `json:"historical_visitor_count"`
	TodayVisitorCount      int64 `json:"today_visitor_count"`
}

type Nav struct {
	MetaData MetaData   `json:"metadata"`
	Children []MetaData `json:"children"`
}
