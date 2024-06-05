package database

import "docsfly/models"

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

func (s *Stack) Push(element LocalMetaCache) {
	s.fileMetas = append(s.fileMetas, element)
}

func (s *Stack) Add(element interface{}) {
	if len(s.fileMetas) == 0 {
		return
	}

	lastLocalMetaCache := &s.fileMetas[len(s.fileMetas)-1]

	switch v := element.(type) {
	case models.Category:
		lastLocalMetaCache.Categorys = append(lastLocalMetaCache.Categorys, v)
	case models.Document:
		lastLocalMetaCache.Documents = append(lastLocalMetaCache.Documents, v)

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
