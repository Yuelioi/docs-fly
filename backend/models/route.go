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

// ResponseBasicData 结构体用于扩展返回数据
type ResponseQueryData struct {
	ClientTime time.Time   `json:"client_time"`
	IP         string      `json:"ip"`
	ServerTime time.Time   `json:"server_time"`
	StatusCode int         `json:"status_code"`
	MaxSize    int         `json:"max_size"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	Data       interface{} `json:"data"`
}
