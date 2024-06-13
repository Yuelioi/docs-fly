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

// 只搜索文件夹
func FindFolders(db *gorm.DB) *gorm.DB {
	return db.Where("is_dir = ?", 1)
}

// 只搜索文件
func FindFiles(db *gorm.DB) *gorm.DB {
	return db.Where("is_dir = ?", 0)
}
