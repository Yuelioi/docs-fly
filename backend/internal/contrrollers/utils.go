package controllers

import (
	"docsfly/internal/models"
	"regexp"
	"strings"
	"sync"

	extractor "github.com/huantt/plaintext-extractor"
	"gorm.io/gorm"
)

var globalCache sync.Map

func getCategoryAndBookByUrl(url string) (string, string) {
	urlList := strings.Split(url, "/")
	if len(urlList) < 2 {
		return "", ""
	}

	return urlList[0], urlList[1]
}

func getFilepathByURL(db *gorm.DB, url string) string {
	var filepath string
	db.Model(models.Entry{}).Where("url = ?", url).Select("filepath").Scan(&filepath)
	return filepath
}

func findClosestDoc(chapter models.Entry, docs []models.Entry) models.Entry {
	var closestDoc models.Entry
	minOrder := int(^uint(0) >> 1) // 初始化为最大值

	for _, doc := range docs {
		if strings.HasPrefix(doc.URL, chapter.URL) && int(doc.Order) < minOrder {
			closestDoc = doc
			minOrder = int(doc.Order)
		}
	}

	return closestDoc
}

// home

// 提取纯文本数据 过滤掉符号
func extractPlainText(markdownContent string) (output *string, err error) {
	extractor := extractor.NewMarkdownExtractor()
	output, err = extractor.PlainText(markdownContent)

	if err != nil {
		return
	}

	var w sync.WaitGroup

	w.Add(1)

	go func() {
		toReplaces := []string{"(\n\\s)+", "-", "\\|", "#", " ", "\t", "\r", "\n", "<iframe[^>]*>.*?</iframe>"}
		for _, toReplace := range toReplaces {
			re := regexp.MustCompile(toReplace)
			*output = re.ReplaceAllString(*output, " ")
		}
		w.Done()
	}()

	w.Wait()
	return output, nil
}

// 根据关键词获取对应索引
func indexOfKeywordInRuneSlice(runeSlice []rune, keyword string) int {
	key := []rune(keyword)

	for idx := range runeSlice {
		// 检查索引范围，确保关键词不会超出切片边界
		if idx+len(key) > len(runeSlice) {
			break
		}

		if string(runeSlice[idx:idx+len(key)]) == keyword {
			return idx
		}
	}
	return -1
}

func filepath2Params(path string) (string, string, string, bool) {
	pathList := strings.Split(path, "/")

	if len(pathList) < 4 {
		return "", "", "", false
	} else {
		return pathList[0], pathList[1], pathList[2], true
	}
}
