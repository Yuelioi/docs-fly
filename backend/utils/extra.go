package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// 将文件路径转换为适合在网站上使用的路径
func ConvertFilepathToURL(filepath string) string {
	// 去除空格
	cleanedPath := strings.ReplaceAll(filepath, " ", "")

	// 定义要替换的特殊字符
	specialChars := regexp.MustCompile(`[!@#$%^&*()+=\[\]{};:'"\\|,<>?]+`)
	cleanedPath = specialChars.ReplaceAllString(cleanedPath, "")

	// 替换连续的多个斜杠为单个斜杠
	cleanedPath = strings.ReplaceAll(cleanedPath, "//", "/")

	return cleanedPath
}

// 反斜杠改为正斜杠
func ReplaceSlash(input string) string {
	return strings.ReplaceAll(input, "\\", "/")
}

func IsZeroType(data interface{}) bool {
	val := reflect.ValueOf(data)
	return reflect.DeepEqual(val.Interface(), reflect.Zero(val.Type()).Interface())
}

func DurationToString(d time.Duration) string {
	seconds := int(d.Seconds()) % 60
	milliseconds := d.Milliseconds() % 1000

	var result string

	if seconds > 0 {
		result += fmt.Sprintf("%d.", seconds)
		result += fmt.Sprintf("%03d秒", milliseconds)
	} else {
		result += fmt.Sprintf("%d毫秒", milliseconds)
	}

	return result
}
