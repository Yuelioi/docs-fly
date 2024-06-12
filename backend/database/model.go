package database

import (
	"docsfly/models"
)

type LocalMetaDatasCache struct {
	ParentFolder string `json:"-"`
	NeedWrite    bool
	Documents    []models.MetaData `json:"documents"`
	Categorys    []models.MetaData `json:"categorys"`
}
