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

var ErrSkip = errors.New("skip")

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
		Cats        map[string]models.Category
		CatsSummary map[string]models.Category
		Docs        map[string]models.Document
		DocsSummary map[string]models.Document
	}
	Local struct {
		Metas   map[string]MetaDatasCache
		Summary map[string]MetaDatasCache
	}
}

func initAdminAccount(db *gorm.DB) {
	var user models.User
	db.Model(models.User{}).Where("username =?", global.AppConfig.Username).Find(&user)

	if user.ID == 0 {
		CreateAdminAccount(db)
	}
}

func initMetaMaps(dbCats []models.Category, dbDocs []models.Document) *MetaMaps {
	metaMaps := MetaMaps{}

	metaMaps.DB.Cats = make(map[string]models.Category)
	metaMaps.DB.CatsSummary = make(map[string]models.Category)
	metaMaps.DB.Docs = make(map[string]models.Document)
	metaMaps.DB.DocsSummary = make(map[string]models.Document)

	for _, cat := range dbCats {
		metaMaps.DB.Cats[cat.Filepath] = cat
		metaMaps.DB.CatsSummary[cat.Filepath] = cat
	}

	for _, doc := range dbDocs {
		metaMaps.DB.Docs[doc.Filepath] = doc
		metaMaps.DB.DocsSummary[doc.Filepath] = doc
	}

	// 本地文件真实元数据映射表
	metaMaps.Local.Metas = make(map[string]MetaDatasCache)

	// 临时元数据映射表, 后续用于写入文件
	metaMaps.Local.Summary = make(map[string]MetaDatasCache)

	return &metaMaps
}

func populateMetaMaps(metaMaps MetaMaps, relative_parent string, parent string) {
	if _, exists := metaMaps.Local.Metas[relative_parent]; !exists {

		localMeta := models.MetaDatas{
			Categorys: make([]models.MetaData, 0), // 初始化为空切片
			Documents: make([]models.MetaData, 0), // 初始化为空切片
		}

		err := utils.ReadJson(filepath.Join(parent, global.AppConfig.MetaFile), &localMeta)

		if err == nil {

			if localMeta.Categorys == nil {
				localMeta.Categorys = make([]models.MetaData, 0)
			}
			if localMeta.Documents == nil {
				localMeta.Documents = make([]models.MetaData, 0)
			}
			metaMaps.Local.Metas[relative_parent] = MetaDatasCache{
				ParentFolder: parent,
				Categorys:    localMeta.Categorys,
				Documents:    localMeta.Documents,
			}
		}
	}
}

func initMetaData(info os.FileInfo, relative_path, relative_parent string) models.MetaData {
	Depth := strings.Count(relative_path, "/")

	return models.MetaData{
		Name:     info.Name(),
		Title:    info.Name(),
		Depth:    Depth,
		Icon:     "",
		Status:   1,
		Filepath: relative_path,
		WebPath:  utils.ConvertFilepathToWebPath(relative_parent),
	}
}

func initLocalMetaCache(metaMaps *MetaMaps, localMetaCache *MetaDatasCache, relative_parent, parent string) {
	// 初始化Summary
	if value, exists := metaMaps.Local.Summary[relative_parent]; !exists {
		// 本地不存在就初始化
		*localMetaCache = MetaDatasCache{
			ParentFolder: parent,
			NeedWrite:    false,
			Categorys:    make([]models.MetaData, 0),
			Documents:    make([]models.MetaData, 0),
		}
		metaMaps.Local.Summary[relative_parent] = *localMetaCache
	} else {
		*localMetaCache = value
	}
}

func initItemMeta(metadata *models.MetaData, localMeta *models.MetaData, dbMeta *models.MetaData, defaultOrder uint) {
	if localMeta != nil {
		metadata.Order = localMeta.Order
		metadata.Icon = localMeta.Icon
		metadata.Status = localMeta.Status
		metadata.Title = localMeta.Title
	} else if dbMeta != nil {
		metadata.Order = dbMeta.Order
		metadata.Icon = dbMeta.Icon
		metadata.Status = dbMeta.Status
		metadata.Title = dbMeta.Title
	} else {
		metadata.Order = defaultOrder
	}
}

/*
初始化数据库
*/
func DBInit(db *gorm.DB) error {
	fmt.Println("初始化数据库准备中...")
	start := time.Now()

	// 如果数据库没有用户 则写入管理员数据
	initAdminAccount(db)

	// 存储各个类目总数据 用于批量写入数据库
	dbDatas := DBDatas{}

	// 最新cat的Id
	var lastCatId uint

	// 数据库映射表
	var dbCats []models.Category
	var dbDocs []models.Document
	db.Find(&dbCats)
	db.Find(&dbDocs)

	db.Model(&models.Category{}).Select("id").Order("id DESC").Limit(1).Scan(&lastCatId)

	// meta映射
	metaMaps := initMetaMaps(dbCats, dbDocs)

	// 上一文件深度

	root := utils.ReplaceSlash(global.AppConfig.Resource)

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

		// 初始化一些路径
		path = utils.ReplaceSlash(path)
		relative_path := utils.ReplaceSlash(strings.ReplaceAll(path, root+"/", ""))

		parent := utils.ReplaceSlash(filepath.Dir(path))
		relative_parent := utils.ReplaceSlash(filepath.Dir(relative_path))

		// 初始化 读取本地文件映射表
		populateMetaMaps(*metaMaps, relative_parent, parent)

		// 本地元数据缓存
		var localMetaCache MetaDatasCache

		initLocalMetaCache(metaMaps, &localMetaCache, relative_parent, parent)

		metaData := initMetaData(info, relative_path, relative_parent)

		if info.IsDir() {

			// 修改时间不变 直接跳过
			if value, exists := metaMaps.DB.CatsSummary[relative_path]; exists {
				if value.ModTime.Equal(info.ModTime()) {
					return filepath.SkipDir
				}
			}

			// 查找本地数据与数据库数据
			localMeta := searchMetaDatasCache(metaMaps.Local.Metas[relative_parent], true, info.Name())
			dbMeta := searchDBCatMetaDatas(dbCats, relative_path)

			// 基于本地数据与数据库数据 更新MetaData值 优先使用本地数据 其次是数据库数据
			initItemMeta(&metaData, localMeta, dbMeta, uint(len(metaMaps.Local.Summary[relative_parent].Categorys))+1)

			cat := models.Category{
				MetaData: metaData,
				ModTime:  info.ModTime(),
			}

			// 检查数据库与本地是否一致
			refresh := !compare(localMeta, dbMeta)

			if refresh {
				localMetaCache.NeedWrite = true
			}

			// 使用本地Meta数据或新的Meta数据
			if localMeta != nil {
				// 本地有就用本地的
				localMetaCache.Categorys = append(localMetaCache.Categorys, *localMeta)

			} else {
				// 本地没有用新的
				newMeta := cat.MetaData
				localMetaCache.Categorys = append(localMetaCache.Categorys, newMeta)
			}

			// 储存数据库数据
			if value, exists := metaMaps.DB.CatsSummary[relative_path]; exists {
				// 删除数据库当前条目
				delete(metaMaps.DB.CatsSummary, relative_path)

				if !value.ModTime.Equal(info.ModTime()) || refresh {
					// 如果数据库有 并且修改时间变化, 则更新
					cat.ModTime = info.ModTime()
					dbDatas.Cats.Updates = append(dbDatas.Cats.Updates, cat)
				}

			} else {
				// 数据库没有 则追加 最新分类数+1
				lastCatId += 1
				dbDatas.Cats.Creates = append(dbDatas.Cats.Creates, cat)
			}

		} else {

			if value, exists := metaMaps.DB.DocsSummary[relative_path]; exists {
				if value.ModTime.Equal(info.ModTime()) {
					return nil
				}
			}

			localMeta := searchMetaDatasCache(metaMaps.Local.Metas[relative_parent], false, info.Name())
			dbMeta := searchDBDocMetaDatas(dbDocs, relative_path)

			initItemMeta(&metaData, localMeta, dbMeta, uint(len(metaMaps.Local.Summary[relative_parent].Documents))+1)

			doc := models.Document{
				MetaData: metaData,
				ModTime:  info.ModTime(),
				Locale:   "",
				Content:  "",
			}

			// 判断数据库与本地是否一致
			refresh := !compare(localMeta, dbMeta)

			if refresh {
				localMetaCache.NeedWrite = true
			}

			// 使用本地Meta数据或新的Meta数据
			if localMeta != nil {
				// 本地有就用本地的
				localMetaCache.Documents = append(localMetaCache.Documents, *localMeta)
			} else {
				// 本地没有用新的
				newMeta := doc.MetaData
				localMetaCache.Documents = append(localMetaCache.Documents, newMeta)
			}

			if value, exists := metaMaps.DB.DocsSummary[relative_path]; exists {
				delete(metaMaps.DB.DocsSummary, relative_path)
				doc.CategoryID = value.CategoryID

				if !value.ModTime.Equal(info.ModTime()) || refresh {
					doc.ModTime = info.ModTime()
					dbDatas.Docs.Updates = append(dbDatas.Docs.Updates, doc)
				}

			} else {

				// 判断使用数据库分类ID 还是使用最新ID
				if value, exists := metaMaps.DB.Cats[relative_parent]; exists {
					doc.CategoryID = value.ID
				} else {
					doc.CategoryID = lastCatId
				}

				dbDatas.Docs.Creates = append(dbDatas.Docs.Creates, doc)
			}

		}

		// 更新缓存数据
		metaMaps.Local.Summary[relative_parent] = localMetaCache

		return nil
	})

	if err != nil {
		return err
	}

	// 更新删除的文件
	for _, cat := range metaMaps.DB.CatsSummary {
		dbDatas.Cats.Deletes = append(dbDatas.Cats.Deletes, cat)
	}
	for _, doc := range metaMaps.DB.DocsSummary {
		dbDatas.Docs.Deletes = append(dbDatas.Docs.Deletes, doc)
	}

	fmt.Println("读取数据用时", time.Since(start))

	start = time.Now()
	WriteMetaData(metaMaps.Local.Summary, false)
	fmt.Println("写入meta数据用时", time.Since(start))

	start = time.Now()
	WriteContentToDocsData(&dbDatas.Docs.Creates, &dbDatas.Docs.Updates)
	fmt.Println("读取内容用时", time.Since(start))

	start = time.Now()

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
