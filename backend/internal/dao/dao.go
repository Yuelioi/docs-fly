package dao

import (
	"docsfly/internal/config"
	"docsfly/pkg/logger"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB
var err error
var ns = schema.NamingStrategy{
	TablePrefix:   "db_",
	SingularTable: false, // 表名后面加s
	NoLowerCase:   false, // 字段转小写
}

func init() {
	// Connect to the database
	Db, err = gorm.Open(sqlite.Open(config.Instance.Database.Resource), &gorm.Config{
		NamingStrategy: ns,
	})

	if err != nil {
		logger.Error(map[string]interface{}{"sqlite connect error": err.Error()})
	}

	if Db.Error != nil {
		logger.Error(map[string]interface{}{"database connect error": Db.Error})
	}

	db, err := Db.DB()
	if err != nil {
		logger.Error(map[string]interface{}{"sql error": err.Error()})
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
}
