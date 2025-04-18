package db

import (
	"docsfly/internal/models"

	"gorm.io/gorm"
)

// 数据库数据管理
type DBDocuments struct {
	// 需要更新的数据集合
	Updates []models.Entry

	// 需要创建的数据集合
	Creates []models.Entry

	// 需要删除的数据集合
	Deletes []models.Entry
}

// Batch 批量处理Creates与Deletes
//
// 参数:
//
//	tx       *gorm.DB  - 数据库事务对象
//	datas    []models.Entry  - 需要批量处理的数据
//	batchSize int      - 每批次处理的数据大小
//	method   string    - 操作方法"Create" / "Delete"
func BatchCD(tx *gorm.DB, datas []models.Entry, batchSize int, method string) (err error) {
	length := len(datas)

	for start := 0; start < length; start += batchSize {
		endIndex := start + batchSize
		if endIndex > length {
			endIndex = length
		}

		batch := datas[start:endIndex]

		if method == "Create" {
			if err = tx.Create(&batch).Error; err != nil {
				return
			}
		} else if method == "Update" {
			for _, b := range batch {
				if err = tx.Save(b).Error; err != nil {
					return
				}

			}
		} else {
			for _, entry := range batch {
				if err = tx.Where("filepath = ?", entry.Filepath).Delete(&entry).Error; err != nil {
					return
				}
			}
		}
	}
	return nil
}

// DBUpdate 处理数据库的批量创建、更新和删除操作
//
// 参数:
//
//	db         *gorm.DB        - 数据库连接对象
//	collection DBDocuments     - 包含要批量创建、更新和删除的数据集合
func DBUpdate(db *gorm.DB, collection DBDocuments) (err error) {

	// 无变化 直接跳过
	if len(collection.Creates) == 0 && len(collection.Updates) == 0 && len(collection.Deletes) == 0 {
		return nil
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	batchSize := 128

	err = tx.Transaction(func(tx *gorm.DB) error {

		if len(collection.Creates) > 0 {

			err = BatchCD(tx, collection.Creates, batchSize, "Create")
			if err != nil {
				return err
			}

		}
		if len(collection.Updates) > 0 {
			// 批量更新
			err = BatchCD(tx, collection.Updates, batchSize, "Update")
			if err != nil {
				return err
			}
			// 使用原生 SQL 批量更新

		}
		if len(collection.Deletes) > 0 {
			err = BatchCD(tx, collection.Deletes, batchSize, "Delete")
			if err != nil {
				return err
			}
		}

		return nil

	})

	return err
}
