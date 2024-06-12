package database

import (
	"docsfly/global"
	"docsfly/models"
	"docsfly/utils"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"gorm.io/gorm"
)

var ErrSkip = errors.New("skip")

// 数据库数据
type DBDatas struct {
	Stores  []models.Entry
	Creates []models.Entry
	Updates []models.Entry
	Deletes []models.Entry
}

// Maps
type MetaMaps struct {
	DB struct {
		// 数据库信息
		Store map[string]models.Entry
		// 数据库信息剩余 用于Delete
		Remain map[string]models.Entry
	}
	Local struct {
		// 本地Meta信息
		Store map[string]MetaDatasCache

		// 剩余Meta信息 用于更新
		Remain map[string]MetaDatasCache
	}
}

// 初始化管理员信息 √
func initAdminAccount(db *gorm.DB) {
	var user models.User
	db.Model(models.User{}).Where("username =?", global.AppConfig.Username).Find(&user)

	if user.ID == 0 {
		CreateAdminAccount(db)
	}
}

// 初始化数据库数据结构 √
//
// @param db *gorm.DB 数据库实例
//
// @return *DBDatas 初始化后的数据结构
func initDBDatas(db *gorm.DB) *DBDatas {
	var files []models.Entry
	db.Find(&files)
	return &DBDatas{
		Stores: files,
	}
}

// 初始化映射表, 包括数据库的以及本地的 √
// 其中本地的赋零值
// 数据库的直接填满
func initMetaMaps(dbDatas DBDatas) *MetaMaps {
	metaMaps := &MetaMaps{}

	metaMaps.DB.Store = make(map[string]models.Entry)
	metaMaps.DB.Remain = make(map[string]models.Entry)

	metaMaps.Local.Store = make(map[string]MetaDatasCache)
	metaMaps.Local.Remain = make(map[string]MetaDatasCache)

	for _, file := range dbDatas.Stores {
		metaMaps.DB.Store[file.Filepath] = file
		metaMaps.DB.Remain[file.Filepath] = file
	}

	return metaMaps
}

// 注入本地读取的Meta数据
func populateMetaMaps(metaMaps *MetaMaps, relative_parent string, parent string) {
	if _, exists := metaMaps.Local.Store[relative_parent]; !exists {

		var localMeta models.MetaDatas

		if err := utils.ReadJson(filepath.Join(parent, global.AppConfig.MetaFile), &localMeta); err != nil {
			localMeta = models.MetaDatas{
				Categorys: make([]models.MetaData, 0),
				Documents: make([]models.MetaData, 0),
			}

		} else {
			if localMeta.Categorys == nil {
				localMeta.Categorys = make([]models.MetaData, 0)
			}
			if localMeta.Documents == nil {
				localMeta.Documents = make([]models.MetaData, 0)
			}
		}

		metaMaps.Local.Store[relative_parent] = MetaDatasCache{
			ParentFolder: parent,
			Categorys:    localMeta.Categorys,
			Documents:    localMeta.Documents,
		}
	}
}

// 创建一个基础Meta数据
func createMetaData(info os.FileInfo, relative_path string) models.MetaData {
	Depth := strings.Count(relative_path, "/")

	return models.MetaData{
		Name:     info.Name(),
		Title:    info.Name(),
		Depth:    Depth,
		Icon:     "",
		Status:   1,
		Filepath: relative_path,
		WebPath:  utils.ConvertFilepathToWebPath(relative_path),
	}
}

// 创建本地临时LocalMetas数据 用于后续修改以及是否保存
func initLocalMetaCache(metaMaps *MetaMaps, relative_parent, parent string) MetaDatasCache {
	if value, exists := metaMaps.Local.Store[relative_parent]; !exists {
		// 本地不存在就初始化
		localMetasCache := MetaDatasCache{
			ParentFolder: parent,
			NeedWrite:    false,
			Categorys:    make([]models.MetaData, 0),
			Documents:    make([]models.MetaData, 0),
		}

		return localMetasCache
	} else {
		return value
	}
}

// 初始化Meta数据
func initItemMeta(metadata *models.MetaData, localMeta *models.MetaData, dbMeta *models.MetaData) {
	if localMeta != nil || dbMeta != nil {
		if value, ok := assignIfNotZero(localMeta.Order, dbMeta.Order); ok {
			metadata.Order = value
		}
		if value, ok := assignIfNotZero(localMeta.Icon, dbMeta.Icon); ok {
			metadata.Icon = value
		}
		if value, ok := assignIfNotZero(localMeta.Status, dbMeta.Status); ok {
			metadata.Status = value
		}
		if value, ok := assignIfNotZero(localMeta.Title, dbMeta.Title); ok {
			metadata.Title = value
		}
	}
}

// assignIfNotZero 赋值辅助函数，处理指针类型
func assignIfNotZero[T comparable](localValue, dbValue T) (T, bool) {
	var zeroValue T
	if !isZeroValue(localValue) {
		return localValue, true
	}
	if !isZeroValue(dbValue) {
		return dbValue, true
	}
	return zeroValue, false
}

// isZeroValue 检查值是否为零值，处理指针类型和非指针类型
func isZeroValue[T comparable](value T) bool {
	// 使用反射来处理指针类型的零值检查
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return v.IsNil()
	default:
		var zeroValue T
		return value == zeroValue
	}
}

/*
初始化数据库
*/
func DBInit(db *gorm.DB) error {
	fmt.Println("初始化数据库准备中...")
	start := time.Now()

	// 文档存放根目录
	var root string
	// 映射表，存储数据库/本地Meta信息。
	var metaMaps *MetaMaps
	// 数据库数据汇总
	var dbDatas *DBDatas

	root = utils.ReplaceSlash(global.AppConfig.Resource)

	// 写入管理员数据
	initAdminAccount(db)

	// 从数据库取数据
	dbDatas = initDBDatas(db)

	// 初始化映射表 填充数据库数据
	metaMaps = initMetaMaps(*dbDatas)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 提前跳过特殊文件与文件夹
		err = WalkSkip(root, info, path)

		if err != nil {
			if errors.Is(err, filepath.SkipDir) {
				return filepath.SkipDir
			}
			if errors.Is(err, ErrSkip) {
				return nil
			}
			return err
		}

		// 元数据
		var metaData models.MetaData

		// 本地元数据
		var localMeta *models.MetaData

		// 数据库元数据
		var dbMeta *models.MetaData

		// 本地元数据表缓存
		var localMetasCache MetaDatasCache

		// 本地元数据表
		var localMetas MetaDatasCache

		// Entry
		var entry models.Entry

		// 初始化一些路径
		path = utils.ReplaceSlash(path)
		relative_path := strings.ReplaceAll(path, root+"/", "")

		parent := utils.ReplaceSlash(filepath.Dir(path))
		relative_parent := utils.ReplaceSlash(filepath.Dir(relative_path))

		// 跳过数据库已经存在的文件文件夹
		if info.IsDir() {
			// 父级文件夹修改时间不变 直接跳过
			if value, exists := metaMaps.DB.Store[relative_parent]; exists {
				if value.ModTime.Equal(info.ModTime()) {
					return filepath.SkipDir
				}
			}
		} else {
			// 父级文件夹修改时间不变 直接跳过
			if value, exists := metaMaps.DB.Store[relative_parent]; exists {
				if value.ModTime.Equal(info.ModTime()) {
					return nil
				}
			}
		}

		// 初始化metaMaps 填充本地数据到Store
		populateMetaMaps(metaMaps, relative_parent, parent)

		// 基于本地数据初始化本地 LocalMeta
		localMetas = initLocalMetaCache(metaMaps, relative_parent, parent)

		// 根据文件信息先初始化一个
		metaData = createMetaData(info, relative_path)

		// 初始化 防止空指针
		dbMeta = &models.MetaData{}
		localMeta = &models.MetaData{}

		// 查找数据库的信息
		searchDBMetaDatas(metaMaps.DB.Store, relative_path, dbMeta)
		// 查找本地数据
		searchMetaDatasCache(metaMaps.Local.Store[relative_parent], info, localMeta)

		// 检查数据库与本地是否一致
		isDifferent := !compare(localMeta, dbMeta)

		if isDifferent {
			localMetasCache.NeedWrite = true
		}

		initItemMeta(&metaData, localMeta, dbMeta)

		entry = models.Entry{
			MetaData: metaData,
			ModTime:  info.ModTime(),
		}

		if info.IsDir() {
			entry.IsDir = true

			if entry.Order == 0 {
				entry.Order = uint(len(metaMaps.Local.Store[relative_parent].Categorys)) + 1
			}

			if localMetasCache.NeedWrite {
				if value, exists := metaMaps.Local.Remain[relative_parent]; exists {
					localMetasCache = value
				} else {
					localMetasCache = MetaDatasCache{}
				}
				localMetasCache.Categorys = append(localMetasCache.Categorys, entry.MetaData)
			}

		} else {

			if entry.Order == 0 {
				entry.Order = uint(len(metaMaps.Local.Store[relative_parent].Documents)) + 1
			}

			localMetasCache.Documents = append(localMetasCache.Documents, entry.MetaData)

		}

		// 数据库有则先删除 如果修改 = 更新, 数据库没有说明要新增
		if value, exists := metaMaps.DB.Store[relative_path]; exists {
			delete(metaMaps.DB.Remain, relative_path)
			fmt.Printf("relative_path: %v\n", relative_path)

			if !value.ModTime.Equal(info.ModTime()) || isDifferent {
				entry.ModTime = info.ModTime()
				dbDatas.Updates = append(dbDatas.Updates, entry)
			}
		} else {
			dbDatas.Creates = append(dbDatas.Creates, entry)
		}

		// 更新缓存数据
		metaMaps.Local.Remain[relative_parent] = localMetasCache

		return nil
	})

	if err != nil {
		return err
	}

	// 数据库要删除的 就是最后还剩下的
	for _, entry := range metaMaps.DB.Remain {
		dbDatas.Deletes = append(dbDatas.Deletes, entry)
	}

	fmt.Println("读取数据用时", time.Since(start))

	start = time.Now()
	WriteMetaData(metaMaps.Local.Remain, false)
	fmt.Println("写入meta数据用时", time.Since(start))

	start = time.Now()
	WriteContentToDocsData(&dbDatas.Creates, &dbDatas.Updates)
	fmt.Println("读取内容用时", time.Since(start))

	start = time.Now()

	collections := DBCollections{
		Creates: []interface{}{dbDatas.Creates},
		Updates: []interface{}{dbDatas.Updates},
		Deletes: []interface{}{dbDatas.Deletes},
		Models:  []interface{}{models.Entry{}},
	}

	err = DBUpdate(db, collections)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	fmt.Println("保存数据库用时", time.Since(start))

	return nil
}
