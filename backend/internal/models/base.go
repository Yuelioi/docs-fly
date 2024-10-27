package models

// 基础信息
type MetaData struct {
	Name     string `json:"name"`     // 文件名
	Title    string `json:"title"`    // 显示名称
	Depth    int    `json:"depth"`    //层级深度
	Order    uint   `json:"order"`    // 所在父级顺序
	IsDir    bool   `json:"is_dir"`   // 是否为文件夹
	Icon     string `json:"icon"`     // 图标
	Status   uint   `json:"status"`   // 状态 0=草稿 1=发布 2=垃圾桶?
	Filepath string `json:"filepath"` // 文件真实路径
	URL      string `json:"url"`      // 文件网站路径
}

// 本地显示信息
type MetaDatas struct {
	Documents  []MetaData `json:"documents"`
	Categories []MetaData `json:"categories"`
}
