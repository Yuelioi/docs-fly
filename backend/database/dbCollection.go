package database

import (
	"reflect"
	"sync"

	"gorm.io/gorm"
)

// 数据库数据管理

type Collections struct {
	Updates []interface{}
	Creates []interface{}
	Deletes []interface{}
	Models  []interface{}
}

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

func DBUpdate(db *gorm.DB, collection Collections) (err error) {

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
