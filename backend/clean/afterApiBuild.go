package main

import (
	"fmt"
	"os"
	"strings"
)

// 生成api后的清理工作

func main() {
	// 定义文件路径
	filePath := "docsfly\\internal\\types\\types.go"

	// 读取文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return
	}

	// 将文件内容按行分割
	lines := strings.Split(string(content), "\n")

	// 初始化变量
	var newLines []string
	foundPackageTypes := false

	for _, line := range lines {
		if !foundPackageTypes && strings.TrimSpace(line) == "package types" {
			// 找到 package types 行，在其下方追加 import 语句
			newLines = append(newLines, line)
			newLines = append(newLines, "")
			newLines = append(newLines, "import \"github.com/lib/pq\"")
			foundPackageTypes = true
		} else {
			// 替换 []string 为 pq.StringArray
			newLine := strings.ReplaceAll(line, "[]string", "pq.StringArray")
			newLines = append(newLines, newLine)
		}
	}

	// 将修改后的内容写回文件
	newContent := strings.Join(newLines, "\n")
	err = os.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		fmt.Printf("Failed to write file: %v\n", err)
		return
	}

	fmt.Println("File updated successfully.")
}
