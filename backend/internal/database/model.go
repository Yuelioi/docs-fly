package database

import (
	"docsfly/models"
)

type LocalMetaDatasCache struct {
	ParentFolder string `json:"-"`
	NeedWrite    bool
	Documents    []models.MetaData `json:"documents"`
	Categories   []models.MetaData `json:"categories"`
}
