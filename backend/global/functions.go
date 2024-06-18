package global

import (
	"bufio"
	"bytes"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func WriteConfigToFile(configName string) error {
	// 打开文件以进行读操作
	file, err := os.OpenFile("conf/config.toml", os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 获取 AppConfig 结构体的反射值
	v := reflect.ValueOf(AppConfig)

	// 读取原文件的内容
	var buffer bytes.Buffer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// 检查是否是待修改字段所在行
		if strings.HasPrefix(line, configName) {
			// 修改字段值
			field := v.FieldByName(configName)
			// 如果字段存在且是可设置的，则修改字段值

			line = configName + " = " + convertToString(field)

		}
		// 将修改后的行写入缓冲区
		buffer.WriteString(line + "\n")
	}

	// 将缓冲区的内容写回文件
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	if _, err := file.Write(buffer.Bytes()); err != nil {
		return err
	}

	// 截断文件，确保只有修改后的内容
	if err := file.Truncate(int64(buffer.Len())); err != nil {
		return err
	}

	return nil
}

// 将字段正确的转为字符串
func convertToString(field reflect.Value) string {
	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(field.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(field.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(field.Float(), 'f', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(field.Bool())
	case reflect.String:
		return field.String()
	default:
		return ""
	}
}
