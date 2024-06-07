package database

import (
	"docsfly/global"
	"docsfly/models"
	"docsfly/utils"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
)

type DBDatas struct {
	Cats struct {
		Creates []models.Category
		Updates []models.Category
		Deletes []models.Category
	}
	Docs struct {
		Creates []models.Document
		Updates []models.Document
		Deletes []models.Document
	}
}

type MetaMaps struct {
	DB struct {
		Cats map[string]models.Category
		Docs map[string]models.Document
	}
	Local struct {
	}
}

/*
初始化数据库
*/
func DBInit(db *gorm.DB) error {
	fmt.Println("初始化数据库准备中...")
	start := time.Now()

	// 如果数据库没有用户 则写入管理员数据
	var user models.User
	db.Model(models.User{}).Where("username =?", global.AppConfig.Username).Find(&user)

	if user.ID == 0 {
		CreateAdminAccount(db)
	}

	// 存储各个类目总数据 用于批量写入数据库
	dbDatas := DBDatas{}

	// 数据库映射表
	var dbCats []models.Category
	var dbDocs []models.Document
	db.Find(&dbCats)
	db.Find(&dbDocs)

	// metaMaps := MetaMaps{}

	dbCatsMap := make(map[string]models.Category)
	dbDocsMap := make(map[string]models.Document)

	for _, cat := range dbCats {
		dbCatsMap[cat.Filepath] = cat
	}

	for _, doc := range dbDocs {
		dbDocsMap[doc.Filepath] = doc
	}

	// 本地文件真实元数据映射表
	localMetas := make(map[string]LocalMetaDatasCache)

	// 临时元数据映射表, 后续用于写入文件
	localMetasCache := make(map[string]LocalMetaDatasCache)

	// 上一文件深度
	var catCount uint = 0

	root := global.AppConfig.Resource

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		err = WalkSkip(root, info, path, err)

		if err != nil {
			if errors.Is(err, filepath.SkipDir) {
				return filepath.SkipDir
			}
			if errors.Is(err, ErrSkip) {
				return nil
			}
			return err
		}

		relative_path := strings.ReplaceAll(path, root+"\\", "")
		Depth := strings.Count(relative_path, "\\")

		parent := filepath.Dir(path)
		relative_parent := filepath.Dir(relative_path)

		metaData := models.MetaData{
			Name:     info.Name(),
			Title:    info.Name(),
			Depth:    Depth,
			Icon:     "",
			Status:   1,
			Filepath: relative_path,
		}

		// 初始化 读取本地文件映射表
		if _, exists := localMetas[relative_parent]; !exists {
			var localMeta models.LocalMetaDatas
			err := utils.ReadJson(parent, localMeta)
			if err == nil {
				localMetasCache[relative_parent] = LocalMetaDatasCache{
					ParentFolder: parent,
					Categorys:    localMeta.Categorys,
					Documents:    localMeta.Documents,
				}
			}
		}

		// 读一下当前分组记录的元数据
		metaCache, exists := localMetasCache[relative_parent]
		if !exists {
			// 不存在就初始化
			metaCache = LocalMetaDatasCache{
				ParentFolder: parent,
				Categorys:    []models.LocalMetaData{},
				Documents:    []models.LocalMetaData{},
			}
		}

		if info.IsDir() {
			catCount++
			cat := models.Category{
				MetaData: metaData,
				ModTime:  info.ModTime(),
				Display:  true,
			}

			localMeta := searchLocalMetaDatasCache(localMetasCache[relative_parent], true, info.Name())

			// 使用本地Meta数据或新的Meta数据
			if localMeta != nil {
				// 本地有就用本地的
				metaCache.Categorys = append(metaCache.Categorys, *localMeta)
			} else {
				// 本地没有用新的
				newMeta := convertMetaData(cat.MetaData)
				metaCache.Categorys = append(metaCache.Categorys, newMeta)
				cat.Order = uint(len(metaCache.Categorys))
			}

			// 更新缓存数据
			localMetasCache[relative_parent] = metaCache

			// 储存数据库数据
			if value, exists := dbCatsMap[relative_path]; exists {
				// 删除数据库当前条目
				delete(dbCatsMap, relative_path)

				if !value.ModTime.Equal(info.ModTime()) {
					// 如果数据库有 并且修改时间变化, 则更新
					dbDatas.Cats.Updates = append(dbDatas.Cats.Updates, cat)
				}

			} else {
				// 数据库没有 则追加
				dbDatas.Cats.Creates = append(dbDatas.Cats.Creates, cat)
			}

		} else {

			doc := models.Document{
				MetaData:   metaData,
				ModTime:    info.ModTime(),
				Locale:     "",
				Content:    "",
				CategoryID: catCount,
			}

			localMeta := searchLocalMetaDatasCache(localMetasCache[relative_parent], false, info.Name())

			// 使用本地Meta数据或新的Meta数据
			if localMeta != nil {
				// 本地有就用本地的
				metaCache.Documents = append(metaCache.Documents, *localMeta)
			} else {
				// 本地没有用新的
				newMeta := convertMetaData(doc.MetaData)
				metaCache.Documents = append(metaCache.Documents, newMeta)
				doc.Order = uint(len(metaCache.Documents))
			}

			if value, exists := dbDocsMap[relative_path]; exists {
				delete(dbDocsMap, relative_path)
				if !value.ModTime.Equal(info.ModTime()) {
					dbDatas.Docs.Updates = append(dbDatas.Docs.Updates, doc)
				}

			} else {
				dbDatas.Docs.Creates = append(dbDatas.Docs.Creates, doc)
			}

		}

		// 更新缓存数据
		localMetasCache[relative_parent] = metaCache

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("读取数据用时", time.Since(start))

	start = time.Now()
	WriteLocalMetaData(localMetasCache)
	fmt.Println("写入meta数据用时", time.Since(start))

	start = time.Now()
	WriteContentToDocsData(&dbDatas.Docs.Creates, &dbDatas.Docs.Updates)
	fmt.Println("读取内容用时", time.Since(start))

	start = time.Now()

	for _, cat := range dbCatsMap {
		dbDatas.Cats.Deletes = append(dbDatas.Cats.Deletes, cat)
	}
	for _, doc := range dbDocsMap {
		dbDatas.Docs.Deletes = append(dbDatas.Docs.Deletes, doc)
	}

	collections := Collections{
		Creates: []interface{}{dbDatas.Cats.Creates, dbDatas.Docs.Creates},
		Updates: []interface{}{dbDatas.Cats.Updates, dbDatas.Docs.Updates},
		Deletes: []interface{}{dbDatas.Cats.Deletes, dbDatas.Docs.Deletes},
		Models:  []interface{}{models.Category{}, models.Document{}},
	}

	err = DBUpdate(db, collections)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	fmt.Println("保存数据库用时", time.Since(start))

	return nil
}
