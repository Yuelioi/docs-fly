package database

import (
	"docsfly/internal/config"
	"docsfly/internal/models"
	"docsfly/pkg/utils"
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
		Store map[string]models.MetaData

		// 已读取的Meta.json relative_path 路径
		Maps map[string]uint

		// 剩余Meta信息 用于更新
		Remain map[string]LocalMetaDatasCache
	}
}

// 初始化管理员信息 √
func initAdminAccount(db *gorm.DB) {
	var user models.User
	db.Model(models.User{}).Where("username =?", config.Instance.Database.Username).Find(&user)

	if user.ID == 0 {
		CreateAdminAccount(db)
	}
}

// 循环前 初始化数据库数据结构 √
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

// 循环前 初始化映射表√
// 本地的赋零值
// 数据库的直接填满
func initMetaMaps(dbDatas DBDatas) *MetaMaps {
	metaMaps := &MetaMaps{}

	metaMaps.DB.Store = make(map[string]models.Entry)
	metaMaps.DB.Remain = make(map[string]models.Entry)

	metaMaps.Local.Store = make(map[string]models.MetaData)
	metaMaps.Local.Remain = make(map[string]LocalMetaDatasCache)
	metaMaps.Local.Maps = make(map[string]uint)

	for _, file := range dbDatas.Stores {
		metaMaps.DB.Store[file.Filepath] = file
		metaMaps.DB.Remain[file.Filepath] = file
	}

	return metaMaps
}

// 初始化本地Store
// 已读取, 直接尝试获取meta数据
// 未读取, 直接填充
// 并且获取本地meta
func initLocalStore(metaMaps *MetaMaps, relative_path, relative_parent, parent string) (localMeta *models.MetaData) {

	var local_MetasCache LocalMetaDatasCache

	if _, exists := metaMaps.Local.Maps[relative_parent]; !exists {

		// 不存在就从本地读取
		if err := utils.ReadJson(filepath.Join(parent, config.Instance.Database.MetaFile), &local_MetasCache); err == nil {
			// 本地有直接写入 但是要防止零值
			if local_MetasCache.Categories != nil {
				for _, cat := range local_MetasCache.Categories {
					metaMaps.Local.Store[cat.Filepath] = cat
				}
			}

			if local_MetasCache.Documents != nil {
				for _, doc := range local_MetasCache.Documents {
					metaMaps.Local.Store[doc.Filepath] = doc
				}
			}
		}

		// 初始化order
		metaMaps.Local.Maps[relative_parent] = 1
	}

	if value, exist := metaMaps.Local.Store[relative_path]; exist {
		localMeta = &value
	} else {
		localMeta = &models.MetaData{}
	}

	return
}

func findLocalMetas(maps *MetaMaps, relative_path string) *LocalMetaDatasCache {
	if value, exist := maps.Local.Remain[relative_path]; exist {
		return &value
	}
	return &LocalMetaDatasCache{
		Documents:  []models.MetaData{},
		Categories: []models.MetaData{},
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
		URL:      utils.ConvertFilepathToURL(relative_path),
	}
}

// 刷新Meta数据
func refreshItemMeta(metadata *models.MetaData, localMeta *models.MetaData, dbMeta *models.MetaData) {
	if localMeta != nil || dbMeta != nil {
		if value, ok := assignIfNotZero(localMeta.Order, dbMeta.Order); ok {
			metadata.Order = value
		}
		if value, ok := assignIfNotZero(localMeta.IsDir, dbMeta.IsDir); ok {
			metadata.IsDir = value
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
		if value, ok := assignIfNotZero(localMeta.URL, dbMeta.URL); ok {
			metadata.URL = value
		}
	}
}

// assignIfNotZero 赋值辅助函数，防止零值
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

	root = utils.ReplaceSlash(config.Instance.Database.Resource)

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

		// 元数据 用于更新
		var metaData models.MetaData

		// 本地元数据 只读
		var localMeta *models.MetaData

		// 数据库元数据 只读
		var dbMeta *models.MetaData

		// 本地当前元数据表缓存 用于更新
		var localMetasCache *LocalMetaDatasCache

		// 初始化一些路径
		path = utils.ReplaceSlash(path)
		relative_path := strings.ReplaceAll(path, root+"/", "")

		parent := utils.ReplaceSlash(filepath.Dir(path))
		relative_parent := utils.ReplaceSlash(filepath.Dir(relative_path))

		// 跳过数据库已经存在的文件夹, 文件夹无法跳过,因为子文件修改, 并不会改变文件夹修改时间
		if !info.IsDir() {
			// 文件修改时间不变 直接跳过
			if value, exists := metaMaps.DB.Store[relative_parent]; exists {
				if value.ModTime.Equal(info.ModTime()) {
					return nil
				}
			}
		}

		// 初始化 不一定有,防止空指针
		dbMeta = &models.MetaData{}
		localMeta = &models.MetaData{}

		// 初始化metaMaps 填充本地数据到Store
		localMeta = initLocalStore(metaMaps, relative_path, relative_parent, parent)

		localMetasCache = findLocalMetas(metaMaps, relative_parent)

		var entry models.Entry

		// 初始化entry 查找数据库 有就直接用数据库的
		if value, exists := metaMaps.DB.Store[relative_path]; exists {
			entry = value
			dbMeta = &entry.MetaData

		} else {
			entry = models.Entry{
				ModTime: info.ModTime(),
			}
		}

		// 查找当前本地元数据
		if value, exists := metaMaps.Local.Store[relative_path]; exists {
			localMeta = &value
		}

		// 根据文件信息先初始化一个
		metaData = createMetaData(info, relative_path)

		// 基于本地 以及数据库的 更新Metadata
		// 可以根据是否有数据库信息 是否有本地信息提前判断
		// 但是常态是本地数据库都有信息,因此忽略
		refreshItemMeta(&metaData, localMeta, dbMeta)

		// 数据库 当前 本地值 三值合一
		entry.MetaData = metaData

		// 检查数据库内容是否变动
		dbNeedUpdate := compare(&metaData, dbMeta)

		// 检查本地数据是否变动
		metaNeedUpdate := compare(&metaData, localMeta)

		if metaNeedUpdate {
			localMetasCache.ParentFolder = relative_parent
			localMetasCache.NeedWrite = true
		}

		if entry.Order == 0 {
			entry.Order = metaMaps.Local.Maps[relative_parent]
		}

		if info.IsDir() {
			entry.IsDir = true
			// 追加真实(本地要用的)分类内容
			localMetasCache.Categories = append(localMetasCache.Categories, entry.MetaData)

		} else {
			// 追加真实文档内容
			localMetasCache.Documents = append(localMetasCache.Documents, entry.MetaData)

		}

		// 数据库有则先删除 如果修改 = 更新, 数据库没有说明要新增
		if value, exists := metaMaps.DB.Store[relative_path]; exists {
			delete(metaMaps.DB.Remain, relative_path)

			// 文件  	数据变动 => 更新
			// 文件夹	文件夹下文件增删变动 (注意: 文件内容变化,不会使父级文件夹修改时间变化)
			if !value.ModTime.Equal(info.ModTime()) || dbNeedUpdate || (info.IsDir() && checkReadme(*metaMaps, relative_path)) {
				fmt.Printf("info.ModTime(): %v\n", info.ModTime())
				fmt.Println(value.ModTime)
				fmt.Println(value.ModTime.Equal(info.ModTime()))

				entry.ModTime = info.ModTime()

				dbDatas.Updates = append(dbDatas.Updates, entry)
			}

		} else {
			dbDatas.Creates = append(dbDatas.Creates, entry)
		}

		// 当前层级的Order+1
		metaMaps.Local.Maps[relative_parent] += 1

		// 储存需要修改的本地metas, 暂时别判断, 可能要全局修改
		metaMaps.Local.Remain[relative_parent] = *localMetasCache

		return nil
	})

	if err != nil {
		return err
	}

	// 数据库要删除的 就是最后还剩下的, 同时还要删除LocalRemain里对应的此条
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
		Creates: dbDatas.Creates,
		Updates: dbDatas.Updates,
		Deletes: dbDatas.Deletes,
	}

	// 第一次初始化会新增meta.json,导致父级文件夹修改时间变化
	err = DBUpdate(db, collections)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	fmt.Println("保存数据库用时", time.Since(start))

	return nil
}
