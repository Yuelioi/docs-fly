package models

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
	MetaData `json:"metadata"`
	Children []MetaData `json:"children"`
}
