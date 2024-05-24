package handlers

import (
	"reflect"
	"regexp"
	"sync"

	extractor "github.com/huantt/plaintext-extractor"
)

func GetFieldValue(obj interface{}, fieldName string) interface{} {
	// 获取对象的反射值
	val := reflect.ValueOf(obj)

	// 如果不是结构体类型，直接返回 nil
	if val.Kind() != reflect.Struct {
		return nil
	}

	// 获取字段的反射值
	fieldVal := val.FieldByName(fieldName)

	// 如果字段不存在，返回 nil
	if !fieldVal.IsValid() {
		return nil
	}

	// 返回字段的值
	return fieldVal.Interface()
}

// 根据关键词获取对应索引
func IndexOfKeywordInRuneSlice(runeSlice []rune, keyword string) int {
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

// 提取纯文本数据 过滤掉符号
func ExtractPlainText(markdownContent string) (output *string, err error) {
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
