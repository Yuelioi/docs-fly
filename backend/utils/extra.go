package utils

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// 查询字符串列表是否在字符串内
func StringsInside(arr []string, query string) bool {
	for _, cur := range arr {
		if strings.Contains(query, cur) {
			return true
		}
	}
	return false
}

// 反斜杠改为正斜杠
func ReplaceSlash(input string) string {
	return strings.ReplaceAll(input, "\\", "/")
}

// 类型转换, 如果失败就用零值
func Transform[T any](data any) T {

	if value, ok := data.(T); ok {
		// 如果转换成功，返回转换后的值
		return value
	}
	// 如果转换失败，使用类型的零值
	var zeroValue T
	return zeroValue
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
