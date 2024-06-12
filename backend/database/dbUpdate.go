package database

import (
	"docsfly/models"
	"sync"

	"gorm.io/gorm"
)

// 数据库数据管理
type DBCollections struct {
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
func Batch(tx *gorm.DB, datas []models.Entry, batchSize int, method string) (err error) {
	length := len(datas)

	for start := 0; start < length; start += batchSize {
		endIndex := start + batchSize
		if endIndex > length {
			endIndex = length
		}

		batch := datas[start:endIndex]

		if method == "Create" {
			if err = tx.Create(&batch).Error; err != nil {
				return err
			}
		} else if method == "Delete" {
			for _, entry := range batch {
				if err = tx.Where("filepath = ?", entry.Filepath).Delete(&entry).Error; err != nil {
					return err
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
//	collection DBCollections     - 包含要批量创建、更新和删除的数据集合
//
// d
// 返回:
//
//	err        error           - 如果操作过程中出现错误，则返回错误，否则返回nil
func DBUpdate(db *gorm.DB, collection DBCollections) (err error) {

	// 无变换 直接跳过
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

			err = Batch(tx, collection.Creates, batchSize, "Create")
			if err != nil {
				return err
			}

		}
		if len(collection.Updates) > 0 {
			var wg sync.WaitGroup
			errCh := make(chan error, len(collection.Updates))
			for _, entry := range collection.Updates {
				wg.Add(1)
				// 筛选出 Cats 和 Docs
				go func(entry models.Entry) {
					defer wg.Done()

					// 基于路径进行更新
					if err := tx.Model(models.Entry{}).Where("filepath = ?", entry.Filepath).Updates(entry).Error; err != nil {
						errCh <- err
						return
					}

				}(entry)
			}
			wg.Wait()
			close(errCh)
			if err := <-errCh; err != nil {
				return err
			}
		}
		if len(collection.Deletes) > 0 {
			err = Batch(tx, collection.Deletes, batchSize, "Delete")
			if err != nil {
				return err
			}
		}

		return nil

	})

	return err
}
