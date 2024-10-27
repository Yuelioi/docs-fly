package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	System    *System
	AppConfig *AppConfig
	DBConfig  *DBConfig
}

type System struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type AppConfig struct {
	AppVersion string `mapstructure:"AppVersion"`
	ApiVersion string `mapstructure:"ApiVersion"`
}

type DBConfig struct {
	Resource string `mapstructure:"resource"`
	MetaFile string `mapstructure:"metafile"`

	Username string `mapstructure:"username"`
	Password string

	LogLevel  string // 添加日志等级配置
	IntroFile string
}

func New() *Config {

	viper.SetConfigName("config") // 配置文件名（不包括扩展名）
	viper.SetConfigType("toml")   // 如果配置文件名中没有扩展名，则需要设置类型
	viper.AddConfigPath(".")      // 查找配置文件所在的路径
	err := viper.ReadInConfig()   // 读取配置文件
	if err != nil {               // 处理错误
		log.Fatalf("Error reading config file, %s", err)
	}

	config := &Config{
		System:    &System{},
		AppConfig: &AppConfig{},
		DBConfig:  &DBConfig{},
	}
	// 解析配置到结构体
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return config
}
