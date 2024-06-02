package database

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
	Init() error
}

type SQLiteManager struct {
	DbFile        string
	IsInitialized bool
}

// NewSQLiteManager 创建一个新的 SQLiteManager 对象
func NewSQLiteManager(dbFile string) *SQLiteManager {
	fmt.Printf("dbFile: %v\n", dbFile)
	sqlManager := SQLiteManager{DbFile: dbFile}
	sqlManager.Init(&sqlManager)
	db, _ := sqlManager.Connect()
	if !sqlManager.IsInitialized {
		DBInit(db)
		sqlManager.IsInitialized = true
	}

	return &sqlManager
}

// 连接数据库
func (m *SQLiteManager) Connect() (*gorm.DB, error) {
	fmt.Println("正在连接数据库...")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second * 3, // 慢速 SQL 阈值
			LogLevel:                  logger.Silent,   // 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略记录器的 ErrRecordNotFound 错误
			ParameterizedQueries:      true,            // 不在 SQL 日志中包含参数
			Colorful:                  false,           // 禁用颜色
		},
	)

	db, _ := gorm.Open(sqlite.Open(m.DbFile), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "db_",
			SingularTable: false, // 表名后面加s
			NoLowerCase:   false, //字段转小写
		},
		Logger: newLogger,
	})

	db.AutoMigrate(
		models.User{},
		models.Category{}, models.Document{},
		models.Visitor{}, models.Comment{},
	)

	return db, nil
}

// 验证文件完整性
func (m *SQLiteManager) Init(manager *SQLiteManager) error {
	fmt.Println("验证数据库文件完整性...")

	if _, err := os.Stat(m.DbFile); os.IsNotExist(err) {

		file, err := os.Create(m.DbFile)

		if err != nil {
			panic("创建文件失败")
		}

		defer file.Close()

		return err
	}

	manager.IsInitialized = true
	return nil
}

var DbManager = NewSQLiteManager(global.AppConfig.Database)
