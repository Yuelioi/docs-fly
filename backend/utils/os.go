package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJson(filepath string, element interface{}) (err error) {
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
