package database

import (
	"docsfly/models"
)

type MetaDatasCache struct {
	ParentFolder string `json:"-"`
	NeedWrite    bool
	Documents    []models.MetaData `json:"documents"`
	Categorys    []models.MetaData `json:"categorys"`
}
