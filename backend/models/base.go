package models

// 基础信息
type MetaData struct {
	Identity    string `json:"identity"`     // 文件名
	DisplayName string `json:"display_name"` // 显示名称
	Order       uint   `json:"order"`        // 所在父级顺序
	Icon        string `json:"icon"`         // 图标
	Hidden      bool   `json:"hidden"`       // 是否隐藏
}
