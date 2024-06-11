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

// Connect 连接 SQLite 数据库并返回 *gorm.DB 对象
//
// 该函数首先检查是否已经存在数据库连接，如果存在则返回现有连接。
// 否则，根据配置的日志级别创建自定义日志记录器，并尝试连接到指定的 SQLite 数据库文件。
// 成功连接后，会执行数据库的自动迁移操作。
//
// 返回:
//
//	*gorm.DB - 成功连接的数据库对象
//	error    - 连接过程中遇到的错误
//
// 日志级别:
//
//	根据 global.AppConfig.LogLevel 配置决定日志级别:
//	  - "silent": logger.Silent
//	  - "error": logger.Error
//	  - "warn": logger.Warn
//	  - "info": logger.Info (默认)
//
// 自定义日志记录器配置:
//   - 慢速 SQL 阈值为 3 秒
//   - 忽略 ErrRecordNotFound 错误
//   - SQL 日志中不包含参数
//   - 启用彩色日志输出
//
// 数据库连接配置:
//   - 表名前缀为 "db_"
//   - 使用复数表名
//   - 字段名不转换为小写
//
// 自动迁移:
//   - 迁移的模型包括: User, Category, Document, Visitor, Comment
//
// 使用示例:
//
//	db, err := m.Connect()
//	if err != nil {
//	  log.Fatalf("数据库连接失败: %v", err)
//	}
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
