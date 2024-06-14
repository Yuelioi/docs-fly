package models

import "time"

// ResponseBasicData 结构体用于扩展返回数据
type ResponseBasicData struct {
	ClientTime time.Time   `json:"client_time"`
	IP         string      `json:"ip"`
	ServerTime time.Time   `json:"server_time"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

type ResponsePageData struct {
	ClientTime time.Time   `json:"client_time"`
	IP         string      `json:"ip"`
	ServerTime time.Time   `json:"server_time"`
	StatusCode int         `json:"status_code"`
	TotalCount int64       `json:"total_count"` // 总记录数
	Page       int         `json:"page"`        // 当前页码 从1开始
	PageSize   int         `json:"page_size"`   // 每页记录数
	Data       interface{} `json:"data"`
}
