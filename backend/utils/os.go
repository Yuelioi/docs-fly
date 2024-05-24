package utils

import (
	"docsfly/global"
	"docsfly/models"
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func GetFilePath(category, book, locale, chapter, section, document string) string {
	return path.Join(global.AppConfig.Resource, category, book, locale, chapter, section, document)
}

func ReadJson(filepath string) (element any, err error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &element)
	if err != nil {
		return
	}
	return
}

func WriteJson(filePath string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return err
	}

	// 写入 JSON 数据到文件
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return err
	}
	return nil
}

func ReadChapter(cat string, docs string) (navs []models.ChapterInfo, err error) {

	jsonFile := fmt.Sprintf("%s\\%s\\%s\\chapter.json", global.AppConfig.Resource, cat, docs)
	mdBytes, err := os.ReadFile(jsonFile)
	if err != nil {
		return
	}
	// 解析 JSON 数据到结构体切片中
	err = json.Unmarshal(mdBytes, &navs)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	return
}

func ReadNav() (element interface{}, err error) {
	jsonFile := fmt.Sprintf("%s\\navs.json", global.AppConfig.Resource)
	element, err = ReadJson(jsonFile)
	if err != nil {
		return
	}
	return
}

func WriteMarkdownFile(category, book, locale, chapter, section, document, content string) error {
	mdFilepath := GetFilePath(category, book, locale, chapter, section, document)
	err := os.WriteFile(mdFilepath, []byte(content), 0644)
	return err
}
