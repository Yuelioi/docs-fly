package database

import (
	"docsfly/internal/global"
	"docsfly/internal/models"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 基于配置文件 用户名与密码创建初始管理员账户
// 参数:
//
//	db *gorm.DB - 数据库连接对象
func CreateAdminAccount(db *gorm.DB) {
	hashedPassword, err := HashPassword(global.AppConfig.DBConfig.Password)
	if err != nil {
		fmt.Println("初始化管理员数据失败")
		return
	}
	userData := models.User{
		Username: global.AppConfig.DBConfig.Username,
		Password: hashedPassword,
	}
	db.Create(&userData)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// WriteContentToDocsData 读取Markdown内容并保存回文档Document{}数据
// 该函数读取给定文档路径下的 Markdown 内容，并将其保存到文档数据的 Content 字段中。
// 文件:直接读取
// 文件夹:读取文件夹下的Readme.md
// 参数:
//
//	datas ...*[]models.Document - 需要处理的文档数据集合
func WriteContentToDocsData(datas ...*[]models.Entry) {

	const maxGoroutines = 500
	guard := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup

	for _, docsDatas := range datas {
		for index, docsData := range *docsDatas {
			wg.Add(1)
			guard <- struct{}{}
			go func(index int, docsData models.Entry) {
				defer wg.Done()

				defer func() { <-guard }()

				var contentPath string
				if docsData.IsDir {
					contentPath = filepath.Join(global.AppConfig.DBConfig.Resource, docsData.Filepath, global.AppConfig.DBConfig.IntroFile)
				} else {
					contentPath = filepath.Join(global.AppConfig.DBConfig.Resource, docsData.Filepath)

				}
				content, err := os.ReadFile(contentPath)
				if err != nil {

					return
				}
				(*docsDatas)[index].Content = string(content)
			}(index, docsData)
		}
	}

	wg.Wait()
}

// compare 比较本地元数据和数据库元数据
// 不相等或者有nil 返回false  相等返回false
func compare(localMeta *models.MetaData, dbMeta *models.MetaData) bool {
	zero := models.MetaData{}

	if localMeta == nil || dbMeta == nil || *localMeta == zero || *dbMeta == zero {
		return true
	}
	return *localMeta != *dbMeta
}

// WriteMetaData 写入本地 meta.json 文件
// 该函数根据需要更新或重写本地的 meta.json 文件。
//
// 参数:
// metas map[string]LocalMetaDatasCache - 需要写入的元数据缓存
// rebuild bool - 是否全部重写（true 表示重写，false 表示只写入修改的部分）
func WriteMetaData(
	metas map[string]LocalMetaDatasCache,
	rebuild bool,
) {

	update_metas := make([]LocalMetaDatasCache, 0)

	for _, meta := range metas {
		if meta.NeedWrite || rebuild {
			update_metas = append(update_metas, meta)
		}

	}

	var wg sync.WaitGroup

	for _, meta := range update_metas {
		wg.Add(1)

		go func(meta LocalMetaDatasCache) {

			defer wg.Done()

			output := models.MetaDatas{
				Categories: meta.Categories,
				Documents:  meta.Documents,
			}

			data, err := json.MarshalIndent(output, "", "    ")
			if err != nil {
				fmt.Println(err.Error())
			}

			outputPath := filepath.Join(global.AppConfig.DBConfig.Resource, meta.ParentFolder, global.AppConfig.DBConfig.MetaFile)

			err = os.WriteFile(outputPath, data, 0644)
			if err != nil {
				fmt.Println(err.Error())
			}

		}(meta)
	}

	wg.Wait()

}

// 检查是否有README 并且README是否变动
// 返回是否需要更新
func checkReadme(metaMaps MetaMaps, relative_path string) bool {

	relative_READMEPath := relative_path + "/" + "README.md"

	realPath := filepath.Join(global.AppConfig.DBConfig.Resource, relative_READMEPath)

	// 检查文件是否存在
	fileInfo, err := os.Stat(realPath)
	if err != nil {
		if os.IsNotExist(err) {
			// README不存在 直接跳过
			return false
		}

	}

	// 获取README当前文件的修改时间
	fileModTime := fileInfo.ModTime()
	// 检查数据库是否存在
	if value, exists := metaMaps.DB.Store[relative_READMEPath]; exists {

		// 比较数据库与当前修改时间, 不相同就修改
		if value.ModTime.Equal(fileModTime) {
			return false
		} else {
			return true

		}
	}
	// 数据库没有,说明是刚创建的,需要修改
	return true

}

// WalkSkip 自定义文件遍历规则
// 如跳过以 "." 和 "_" 开头的目录和文件。
//
// 参数:
// root string - 根目录
// info os.FileInfo - 文件信息
// path string - 文件路径
// err error - 遍历过程中遇到的错误
//
// 返回:
// error - 如果符合跳过条件，则返回 ErrSkip，否则返回 nil
func WalkSkip(root string, info os.FileInfo, path string) error {

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

	if !info.IsDir() && !strings.HasSuffix(strings.ToLower(info.Name()), ".md") {
		return ErrSkip
	}

	if info.Name() == global.AppConfig.DBConfig.MetaFile {
		return ErrSkip
	}
	if info.Name() == "main.db" {
		return ErrSkip
	}

	return nil
}
