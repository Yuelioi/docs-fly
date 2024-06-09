package utils

import "strings"

// 分类 书籍 语言 文档 是否成功
func Filepath2Params(path string) (string, string, string, bool) {
	pathList := strings.Split(path, "/")

	if len(pathList) < 4 {
		return "", "", "", false
	} else {
		return pathList[0], pathList[1], pathList[2], true
	}
}
