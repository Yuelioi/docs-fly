package database

import (
	"reflect"

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
			for _, item := range collection.Updates {
				itemVal := reflect.ValueOf(item)
				itemType := itemVal.Type()

				for _, model := range collection.Models {
					modelVal := reflect.ValueOf(model)
					modelType := modelVal.Type()

					if itemType == modelType {
						if err := tx.Model(model).Where("id = ?", itemVal.FieldByName("ID").Interface()).Updates(item).Error; err != nil {
							return err
						}
						break
					}
				}
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
