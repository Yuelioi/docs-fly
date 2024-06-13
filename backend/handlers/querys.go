package handlers

import (
	"gorm.io/gorm"
)

// 分类层级:0
func FindCategory(db *gorm.DB) *gorm.DB {
	return db.Where("depth = ?", 0)
}

// 书籍层级:1
func FindBook(db *gorm.DB) *gorm.DB {
	return db.Where("depth = ?", 1)
}

// 语言层级:2
func FindLocale(db *gorm.DB) *gorm.DB {
	return db.Where("depth = ?", 2)
}

// 章节层级:3
func FindChapter(db *gorm.DB) *gorm.DB {
	return db.Where("depth = ?", 3)
}

// 只搜索文件0
func FindFile(db *gorm.DB) *gorm.DB {
	return db.Where("is_dir = ?", 0)
}

// 只搜索文件夹1
func FindFolder(db *gorm.DB) *gorm.DB {
	return db.Where("is_dir = ?", 1)
}

// 匹配文件夹前缀,已经补了正斜杠 /
func HasPrefixPath(path string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("filepath Like ?", path+"/%")
	}
}
