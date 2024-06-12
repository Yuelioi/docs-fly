package database

import (
	"reflect"
	"sync"

	"gorm.io/gorm"
)

// 数据库数据管理
type DBCollections struct {
	// 需要更新的数据集合
	Updates []interface{}

	// 需要创建的数据集合
	Creates []interface{}

	// 需要删除的数据集合
	Deletes []interface{}

	// 数据模型, 用于更新数据集合
	Models []interface{}
}

// Batch 批量处理Creates与Deletes
//
// 参数:
//
//	tx       *gorm.DB  - 数据库事务对象
//	datas    []interface{}  - 需要批量处理的数据
//	batchSize int      - 每批次处理的数据大小
//	method   string    - 操作方法"Create" / "Delete"
func Batch(tx *gorm.DB, datas []interface{}, batchSize int, method string) (err error) {
	for _, groups := range datas {
		group := reflect.ValueOf(groups)
		length := group.Len()

		filepaths := make([]string, length)
		for i := 0; i < length; i++ {
			item := group.Index(i).Interface()
			filepath := reflect.ValueOf(item).FieldByName("Filepath").String()
			filepaths[i] = filepath
		}

		for start := 0; start < length; start += batchSize {
			endIndex := start + batchSize
			if endIndex > length {
				endIndex = length
			}
			batch := group.Slice(start, endIndex).Interface()
			if method == "Create" {
				if err = tx.Create(batch).Error; err != nil {
					return err
				}
			} else {

				if err = tx.Where("filepath IN ? ", filepaths).Delete(batch).Error; err != nil {
					return err
				}
			}

		}
	}
	return
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
			for _, update := range collection.Updates {
				wg.Add(1)
				// 筛选出 Cats 和 Docs
				go func(update interface{}) {
					defer wg.Done()
					updateVal := reflect.ValueOf(update)
					switch updateVal.Kind() {
					case reflect.Slice:
						for i := 0; i < updateVal.Len(); i++ {
							item := updateVal.Index(i)
							for _, model := range collection.Models {
								modelVal := reflect.ValueOf(model)
								modelType := modelVal.Type()
								itemType := item.Type()
								filepathValue := item.FieldByName("Filepath").Interface()
								updateFields := item.Interface()

								// 基于路径进行更新

								if itemType == modelType {
									if err := tx.Model(model).Where("filepath = ?", filepathValue).Updates(updateFields).Error; err != nil {
										errCh <- err
										return
									}
									break
								}
							}
						}
					}
				}(update)
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
