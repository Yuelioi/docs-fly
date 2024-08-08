package book

import "docsfly/models"

type BookData struct {
	Url      string          `json:"url"`
	IsDir    bool            `json:"is_dir"`
	MetaData models.MetaData `json:"metadata"`
}
