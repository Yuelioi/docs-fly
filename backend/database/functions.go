package database

import (
	"docsfly/global"
	"docsfly/models"
	"docsfly/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gorm.io/gorm"
)

// 基于app.toml用户名与密码创建初始管理员账户
// 参数:
//
//	db *gorm.DB - 数据库连接对象
func CreateAdminAccount(db *gorm.DB) {
	hashedPassword, err := utils.HashPassword(global.AppConfig.Password)
	if err != nil {
		fmt.Println("初始化管理员数据失败")
		return
	}
	userData := models.User{
		Username: global.AppConfig.Username,
		Password: hashedPassword,
	}
	db.Create(&userData)
}

// WriteContentToDocsData 读取Markdown内容并保存回文档Document{}数据
// 该函数读取给定文档路径下的 Markdown 内容，并将其保存到文档数据的 Content 字段中。
// 参数:
//
//	datas ...*[]models.Document - 需要处理的文档数据集合
func WriteContentToDocsData(datas ...*[]models.Document) {

	const maxGoroutines = 500
	guard := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup

	for _, docsDatas := range datas {
		for index, docsData := range *docsDatas {
			wg.Add(1)
			guard <- struct{}{}
			go func(index int, docsData models.Document) {
				defer wg.Done()

				defer func() { <-guard }()

				docsPath := filepath.Join(global.AppConfig.Resource, docsData.MetaData.Filepath)
				content, err := os.ReadFile(docsPath)
				if err != nil {
					log.Printf("%s", err)
					return
				}
				(*docsDatas)[index].Content = string(content)
			}(index, docsData)
		}
	}

	wg.Wait()
}

// 在缓存中搜索元数据
// 该函数在给定的元数据缓存中根据名称搜索类别或文档的元数据。
//
// 参数:
//
//	cache MetaDatasCache - 元数据缓存
//	isCat bool - 是否为类别（true 表示类别，false 表示文档）
//	name string - 要搜索的元数据名称
//
// 返回:
//
//	*models.MetaData - 找到的元数据指针，如果未找到则返回 nil
func searchMetaDatasCache(cache MetaDatasCache, isCat bool, name string) *models.MetaData {
	if isCat {
		for _, meta := range cache.Categorys {
			if meta.Name == name {
				return &meta
			}
		}
	} else {
		for _, meta := range cache.Documents {
			if meta.Name == name {
				return &meta
			}
		}
	}
	return nil
}

// 在数据库类别缓存中搜索元数据
// 该函数在数据库类别缓存中根据路径搜索类别的元数据
//
// 参数:
//
// cache []models.Category - 类别缓存
// path string - 要搜索的类别路径
// 返回:
// *models.MetaData - 找到的元数据指针，如果未找到则返回 nil
func searchDBCatMetaDatas(cache []models.Category, path string) *models.MetaData {
	for _, item := range cache {
		if item.Filepath == path {
			return &item.MetaData
		}
	}
	return nil
}

// 在数据库文档缓存中搜索元数据
// 该函数在数据库文档缓存中根据路径搜索文档的元数据。

// 参数:
//
// cache []models.Document - 文档缓存
// path string - 要搜索的文档路径
//
// 返回:
// *models.MetaData - 找到的元数据指针，如果未找到则返回 nil
func searchDBDocMetaDatas(cache []models.Document, path string) *models.MetaData {
	for _, item := range cache {
		if item.Filepath == path {
			return &item.MetaData
		}
	}
	return nil
}

// compare 比较本地元数据和数据库元数据
// 该函数比较本地元数据和数据库元数据，判断它们是否相等。
//
// 参数:
// localMeta *models.MetaData - 本地元数据
// dbMeta *models.MetaData - 数据库元数据
//
// 返回:
// bool - 如果两个元数据相等则返回 true，否则返回 false
func compare(localMeta *models.MetaData, dbMeta *models.MetaData) bool {
	if localMeta == nil || dbMeta == nil {
		return false
	}
	return *localMeta == *dbMeta
}

// WriteMetaData 写入本地 meta.json 文件
// 该函数根据需要更新或重写本地的 meta.json 文件。
//
// 参数:
// metas map[string]MetaDatasCache - 需要写入的元数据缓存
// rebuild bool - 是否全部重写（true 表示重写，false 表示只写入修改的部分）
func WriteMetaData(
	metas map[string]MetaDatasCache,
	rebuild bool,
) {

	update_metas := make([]MetaDatasCache, 0)

	for _, meta := range metas {
		if meta.NeedWrite || rebuild {
			update_metas = append(update_metas, meta)
		}

	}

	var wg sync.WaitGroup

	for _, meta := range update_metas {
		wg.Add(1)

		go func(meta MetaDatasCache) {

			defer wg.Done()

			output := models.MetaDatas{
				Categorys: meta.Categorys,
				Documents: meta.Documents,
			}

			data, _ := json.MarshalIndent(output, "", "    ")
			outputPath := filepath.Join(meta.ParentFolder, global.AppConfig.MetaFile)

			os.WriteFile(outputPath, data, 0644)

		}(meta)
	}

	wg.Wait()

}

// WalkSkip 自定义文件遍历规则
// 该函数定义了文件遍历时跳过的规则，比如跳过以 "." 和 "_" 开头的目录和文件。
//
// 参数:
// root string - 根目录
// info os.FileInfo - 文件信息
// path string - 文件路径
// err error - 遍历过程中遇到的错误
//
// 返回:
// error - 如果符合跳过条件，则返回 ErrSkip，否则返回 nil
func WalkSkip(root string, info os.FileInfo, path string, err error) error {

	if path == root {
		return ErrSkip
	}

	if info.IsDir() && strings.HasPrefix(info.Name(), ".") {
		return filepath.SkipDir
	}

	if info.IsDir() && strings.HasPrefix(info.Name(), "_") {
		return filepath.SkipDir
	}

	// if info.IsDir() && strings.ToLower(info.Name()) == "ue" {
	// 	return filepath.SkipDir
	// }

	if !info.IsDir() && strings.HasPrefix(info.Name(), "_") {
		return ErrSkip
	}

	if !info.IsDir() && !strings.HasSuffix(info.Name(), ".md") {
		return ErrSkip
	}

	if info.Name() == global.AppConfig.MetaFile {
		return ErrSkip
	}
	if info.Name() == "main.db" {
		return ErrSkip
	}

	return nil
}
