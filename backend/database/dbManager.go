package database

// 数据库管理

import (
	"docsfly/global"
	"docsfly/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DBManager interface {
	Connect() (*gorm.DB, error)
	Init(*gorm.DB) error
}

type SQLiteManager struct {
	DbFile string
	Db     *gorm.DB
}

func NewSQLiteManager(dbFile string) *SQLiteManager {
	sqlManager := SQLiteManager{DbFile: dbFile}

	db, _ := sqlManager.Connect()
	sqlManager.Init(db)

	return &sqlManager
}

func (m *SQLiteManager) Init(db *gorm.DB) error {
	return DBInit(db)
}

// 连接数据库
func (m *SQLiteManager) Connect() (*gorm.DB, error) {
	fmt.Println("正在连接数据库...")

	if m.Db != nil {
		fmt.Println("使用现有数据库")
		return m.Db, nil
	}

	var logLevel logger.LogLevel

	switch global.AppConfig.LogLevel {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		logLevel = logger.Info // 默认使用 Info 级别
	}

	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second * 3, // 慢速 SQL 阈值
			LogLevel:                  logLevel,        // 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略记录器的 ErrRecordNotFound 错误
			ParameterizedQueries:      true,            // 不在 SQL 日志中包含参数
			Colorful:                  true,            // 禁用颜色
		},
	)

	db, err := gorm.Open(sqlite.Open(m.DbFile), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "db_",
			SingularTable: false, // 表名后面加s
			NoLowerCase:   false, //字段转小写
		},
		Logger: customLogger,
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(
		models.User{},
		models.Category{}, models.Document{},
		models.Visitor{}, models.Comment{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	m.Db = db
	return db, nil
}

var DbManager = NewSQLiteManager(global.AppConfig.Database)
