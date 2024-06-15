package global

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Resource string
	Database string
	MetaFile string

	Addr string

	Username string
	Password string

	LogLevel  string // 添加日志等级配置
	IntroFile string

	AppVersion string
	APIVersion string // API版本
}

var AppConfig Config

func init() {
	// 解析配置文件
	if _, err := toml.DecodeFile("conf/app.toml", &AppConfig); err != nil {
		log.Fatal(err)
	}

}
