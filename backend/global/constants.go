package global

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Resource string
	Database string
	MetaFile string

	Addr   string
	DBMode int

	Username string
	Password string

	LogLevel            string // 添加日志等级配置
	FolderIntroFileName string
}

var AppConfig Config

func init() {
	// 解析配置文件
	if _, err := toml.DecodeFile("conf/app.toml", &AppConfig); err != nil {
		log.Fatal(err)
	}

}
