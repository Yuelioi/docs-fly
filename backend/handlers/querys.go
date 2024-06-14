package handlers

import (
	"docsfly/models"

	"gorm.io/gorm"
)

// Model Entry
func BasicModel(db *gorm.DB) *gorm.DB {
	return db.Model(models.Entry{})
}

// 分类 深度 :0
func FindCategory(db *gorm.DB) *gorm.DB {
	return db.Where("depth = ?", 0)
}

// 书籍 深度:1
func FindBook(db *gorm.DB) *gorm.DB {
	return db.Where("depth = ?", 1)
}

// 语言 深度:2
func FindLocale(db *gorm.DB) *gorm.DB {
	return db.Where("depth = ?", 2)
}

// 章节 深度:3
func FindChapter(db *gorm.DB) *gorm.DB {
	return db.Where("depth = ?", 3)
}

// 文件 only
func FindFile(db *gorm.DB) *gorm.DB {
	return db.Where("is_dir = ?", 0)
}

// Dir only
func FindFolder(db *gorm.DB) *gorm.DB {
	return db.Where("is_dir = ?", 1)
}

// 匹配文件夹前缀,已经补了正斜杠 /
func HasPrefixPath(path string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("filepath Like ?", path+"/%")
	}
}
