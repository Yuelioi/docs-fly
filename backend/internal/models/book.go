package models

type BookData struct {
	Url      string   `json:"url"`
	IsDir    bool     `json:"is_dir"`
	MetaData MetaData `json:"metadata"`
}

type PageData struct {
	TotalCount int64         `json:"total_count"` // 总记录数
	Page       int           `json:"page"`        // 当前页码 从1开始
	PageSize   int           `json:"page_size"`   // 每页记录数
	Pages      []interface{} `json:"pages"`
}
