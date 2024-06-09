package models

type BookData struct {
	Url         string `json:"url"`
	ChapterType string `json:"chapter_type"`
	MetaData    `json:"metadata"`
}
