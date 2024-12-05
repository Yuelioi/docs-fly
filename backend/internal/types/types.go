// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

import "github.com/lib/pq"

type Category struct {
	Model
	CID       string `json:"cid,optional" gorm:"column:cid"`
	Depth     int    `json:"depth,optional,omitempty" yaml:",omitempty"`
	Title     string `json:"title"`
	FullTitle string `json:"full_title"`
	Order     int    `json:"order,optional"`
	Path      string `json:"path,optional,omitempty" gorm:"column:path;unique"`
}

type Document struct {
	Model
	CID      string `json:"cid,optional" gorm:"column:cid"`
	Depth    int    `json:"depth,optional,omitempty" yaml:",omitempty"`
	Title    string `json:"title"`           // 显示名称
	Depth    int    `json:"-"`               // 层级深度
	Order    uint   `json:"order,optional"`  // 所在父级顺序
	Filepath string `json:"filepath;unique"` // 文件真实路径
	Content  string
	URL      string `json:"url"`    // 文件网站路径
	Status   uint   `json:"status"` // 状态 0=草稿 1=发布 2=垃圾桶?
}

type IDResponse struct {
	ID string `json:"id"`
}

type Model struct {
	ID        uint   `json:"-" gorm:"primarykey"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"update_at"`
}