package database

import "docsfly/models"

// 联合的文件/文件夹信息
type FileInfo struct {
	Depth    int
	FileType string
	Document models.Document
	Category models.Category
}

type LocalMetaCache struct {
	Depth     int               `json:"-"`
	Folder    string            `json:"-"`
	Size      int               `json:"-"`
	Hash      string            `json:"-"`
	Documents []models.Document `json:"documents"`
	Categorys []models.Category `json:"categorys"`
}

type Stack struct {
	fileMetas []LocalMetaCache
}

func newStack() *Stack {
	return &Stack{fileMetas: []LocalMetaCache{}}
}

func (s *Stack) Push(element LocalMetaCache) {
	s.fileMetas = append(s.fileMetas, element)
}

func (s *Stack) Add(element FileInfo) {
	if len(s.fileMetas) == 0 {
		return
	}

	lastLocalMetaCache := &s.fileMetas[len(s.fileMetas)-1]

	if element.FileType == "category" {
		lastLocalMetaCache.Categorys = append(lastLocalMetaCache.Categorys, element.Category)
	} else if element.FileType == "document" {
		lastLocalMetaCache.Documents = append(lastLocalMetaCache.Documents, element.Document)
	}

}

func (s *Stack) Pop() *LocalMetaCache {
	if len(s.fileMetas) == 0 {
		return nil
	}
	element := &s.fileMetas[len(s.fileMetas)-1]
	s.fileMetas = s.fileMetas[:len(s.fileMetas)-1]
	return element
}
