package models

// 基础信息
type MetaData struct {
	Name     string `json:"name"`                                                                     // 文件名
	Title    string `json:"title"`                                                                    // 显示名称
	Depth    int    `json:"depth"`                                                                    //层级深度
	Order    uint   `json:"order"`                                                                    // 所在父级顺序
	Icon     string `json:"icon"`                                                                     // 图标
	Status   uint   `json:"status"`                                                                   // 状态 0=草稿 1=发布 2=垃圾桶?
	Filepath string `json:"filepath" gorm:"uniqueIndex:idx_unique_filepath,where:deleted_at IS NULL"` // 所在父级顺序
}

type LocalMetaData struct {
	Name   string `json:"name"`   // 文件名
	Title  string `json:"title"`  // 显示名称
	Order  uint   `json:"order"`  // 所在父级顺序
	Icon   string `json:"icon"`   // 图标
	Status uint   `json:"status"` // 状态 0=草稿 1=发布 2=垃圾桶
}

// 本地显示信息
type LocalMetaDatas struct {
	Documents []LocalMetaData `json:"documents"`
	Categorys []LocalMetaData `json:"categorys"`
}
