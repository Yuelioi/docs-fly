package database

import "docsfly/models"

type LocalMetaDatasCache struct {
	ParentFolder string            `json:"-"`
	Documents    []models.Document `json:"documents"`
	Categorys    []models.Category `json:"categorys"`
}

type Stack struct {
	fileMetas []LocalMetaDatasCache
}

func (s *Stack) Push(item LocalMetaDatasCache) {
	s.fileMetas = append(s.fileMetas, item)
}

func (s *Stack) Add(item interface{}) {
	if len(s.fileMetas) == 0 {
		return
	}

	lastLocalMetaDatasCache := &s.fileMetas[len(s.fileMetas)-1]

	switch v := item.(type) {
	case models.Category:
		lastLocalMetaDatasCache.Categorys = append(lastLocalMetaDatasCache.Categorys, v)
	case models.Document:
		lastLocalMetaDatasCache.Documents = append(lastLocalMetaDatasCache.Documents, v)

	}

}

func (s *Stack) Pop() *LocalMetaDatasCache {
	if len(s.fileMetas) == 0 {
		return nil
	}
	item := &s.fileMetas[len(s.fileMetas)-1]
	s.fileMetas = s.fileMetas[:len(s.fileMetas)-1]
	return item
}
