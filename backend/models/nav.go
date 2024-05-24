package models

// 顶部导航栏数据
type NavData struct {
	MetaData
	Children []MetaData `json:"children"`
}
