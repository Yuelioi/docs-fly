package home

import (
	"regexp"
	"strings"
	"sync"

	extractor "github.com/huantt/plaintext-extractor"
)

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
