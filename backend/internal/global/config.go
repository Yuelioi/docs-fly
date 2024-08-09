package global

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var AppConfig = initConfig()

type Config struct {
	viper        *viper.Viper
	ServerConfig *ServerConfig
	AppConfig    *appConfig
	DBConfig     *DBConfig
}

type ServerConfig struct {
	Addr string
}
type appConfig struct {
	AppVersion string
	ApiVersion string // API版本
}
type DBConfig struct {
	Resource string
	Database string
	MetaFile string

	Username string
	Password string

	LogLevel  string // 添加日志等级配置
	IntroFile string
}

func initConfig() *Config {

	viper := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	dir, err := os.Getwd()

	fmt.Printf("dir: %v\n", dir)
	if err != nil {
		panic(fmt.Errorf("fatal error getwd: %w", err))
	}

	viper.AddConfigPath(filepath.Join(filepath.Dir(filepath.Dir(dir)), "configs"))

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	conf := &Config{viper: viper}
	conf.ReadServerConfig()
	conf.ReadAppConfig()
	conf.ReadDBConfig()

	return conf
}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Addr = c.viper.GetString("Addr")
	c.ServerConfig = sc
}
func (c *Config) ReadAppConfig() {
	ac := &appConfig{}
	ac.AppVersion = c.viper.GetString("AppVersion")
	ac.ApiVersion = c.viper.GetString("ApiVersion")
	c.AppConfig = ac
}
func (c *Config) ReadDBConfig() {
	dc := &DBConfig{}
	dc.Database = c.viper.GetString("Database")
	dc.IntroFile = c.viper.GetString("IntroFile")
	dc.LogLevel = c.viper.GetString("LogLevel")
	dc.MetaFile = c.viper.GetString("MetFile")
	dc.Resource = c.viper.GetString("Resource")
	dc.Username = c.viper.GetString("Username")
	dc.Password = c.viper.GetString("Password")

	c.DBConfig = dc
}
