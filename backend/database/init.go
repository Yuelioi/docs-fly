package database

import (
	"docsfly/global"
	"docsfly/models"
	"docsfly/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
)

/*
初始化数据库: 仅在无数据库, 或者手动修改meta信息后重置数据库

	@params:root
		markdown文件存放位置
	@params:Mode
	1: init模式,第一次创建数据库, 会创建meta信息
	2: update模式, 会按照目录下的meta.json读取设置
*/
func DBInit(db *gorm.DB) error {
	// TODO 初始化数据库没有文件 会验证2次,如果一直失败可能会死循环
	fmt.Println("初始化数据库准备中...")

	root := global.AppConfig.Resource
	Mode := global.AppConfig.DBMode

	if Mode == 0 {
		println("注意!!!当前为初始化模式, 并且会修改文件内容")
	} else if Mode <= 1 {
		println("当前为初始化模式,只会生成meta.json")

	} else if Mode == 2 {
		println("当前为更新模式: 会读取meta信息")
	}

	start := time.Now()

	// 写入管理员数据
	CreateAdminAccount(db)

	// 存储各个类目总数据 用于写入数据库
	catDatas := make([]models.Category, 0)
	bookDatas := make([]models.Book, 0)
	chapterDatas := make([]models.Chapter, 0)
	sectionDatas := make([]models.Section, 0)
	docsDatas := make([]models.Document, 0)

	// 语言
	locale := "zh"

	// 类目索引
	catId, bookId, chapterId, sectionId, docsId := uint(0), uint(0), uint(0), uint(0), uint(0)

	// 类目所在父级排序
	bookOrder, chapterOrder, sectionOrder, docsOrder := uint(0), uint(0), uint(0), uint(0)

	// 类目当前配置,创建一下,防止重复获取
	var (
		currentCatsMeta     *[]models.MetaData
		currentBooksMeta    *[]models.MetaData
		currentSectionsMeta *[]models.MetaData
		currentChaptersMeta *[]models.MetaData
	)

	// 当前子目信息 用于创建本地数据
	currentCatBooks := make([]models.Book, 0)
	currentBookChapters := make([]models.Chapter, 0)

	// 本地meta信息汇总
	summaryMeta := make([]models.MetaData, 0)
	cateLocalMeta := make([]models.MetaDataLocal, 0)
	bookLocalMeta := make([]models.MetaDataLocal, 0)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Println("Error Walk:", err)
			return err
		}
		// 根目录 直接跳过
		if path == root {
			return nil
		}

		// 排除.和_开头的文件
		if strings.HasPrefix(info.Name(), ".") || strings.HasPrefix(info.Name(), "_") {
			return nil
		}

		// 获取文件深度
		// 1=分类层级  2=书籍层级 3语言  4=章节/文档  5=文档
		pathList := strings.Split(strings.TrimPrefix(path, root+"\\"), "\\")
		depth := len(pathList)

		if info.IsDir() {

			if depth == 1 {
				// 分类层级
				// 增加分类索引 重置书籍排序/当前书籍meta信息
				catId += 1
				bookOrder = 0
				currentBooksMeta = nil

				if Mode <= 1 {
					// 当前分类书籍有信息(不是第一次循环) 追加分类信息到本地分类元数据记录 cateLocalMeta
					if len(currentCatBooks) > 0 {
						metas := make([]models.MetaData, 0)
						for _, data := range currentCatBooks {
							metas = append(metas, data.MetaData)
						}

						cateLocalMeta = append(cateLocalMeta, models.MetaDataLocal{
							MetaDatas: metas,
							Filepath:  currentCatBooks[0].Filepath,
						})
						// 加完重置下
						currentCatBooks = nil
					}
				}

				// 更新模式 读取meta.json
				meta := *utils.CreateMeta(info, catId)
				if Mode == 2 {
					currentCatsMeta, meta = CreateMetaByCurrentMetas(currentCatsMeta, path, info, catId)
				}

				// 初始化分类信息
				catData := models.Category{
					MetaData: meta,
					ModTime:  info.ModTime(),
					Filepath: path,
				}

				catDatas = append(catDatas, catData)

				// 增加所有目录信息
				summaryMeta = append(summaryMeta, catData.MetaData)

			} else if depth == 2 {
				// 书籍层级
				bookId += 1
				bookOrder += 1

				meta := *utils.CreateMeta(info, bookOrder)
				if Mode == 2 {
					currentBooksMeta, meta = CreateMetaByCurrentMetas(currentBooksMeta, path, info, bookOrder)
				}

				bookData := models.Book{
					MetaData:   meta,
					ModTime:    info.ModTime(),
					Filepath:   path,
					CategoryID: catId,
				}

				currentCatBooks = append(currentCatBooks, bookData)
				bookDatas = append(bookDatas, bookData)

			} else if depth == 3 {
				// 语言版本
				locale = info.Name()
				chapterOrder = 0
				sectionOrder = 0
				currentChaptersMeta = nil

				if Mode <= 1 {
					if len(currentBookChapters) > 0 {
						metas := make([]models.MetaData, 0)
						for _, data := range currentBookChapters {
							metas = append(metas, data.MetaData)
						}

						bookLocalMeta = append(bookLocalMeta, models.MetaDataLocal{
							MetaDatas: metas,
							Filepath:  currentBookChapters[0].Filepath,
						})
						currentBookChapters = nil
					}
				}

			} else if depth == 4 {
				// 目录大纲Chapter
				chapterId += 1
				chapterOrder += 1
				docsOrder = 0

				meta := *utils.CreateMeta(info, chapterOrder)
				if Mode == 2 {
					currentChaptersMeta, meta = CreateMetaByCurrentMetas(currentChaptersMeta, path, info, chapterOrder)
				}

				chapterData := models.Chapter{
					MetaData:   meta,
					ModTime:    info.ModTime(),
					Filepath:   path,
					Locale:     locale,
					CategoryID: catId,
					BookID:     bookId,
				}

				chapterDatas = append(chapterDatas, chapterData)
				currentBookChapters = append(currentBookChapters, chapterData)

			} else if depth == 5 {
				// 小节 Section(傻逼ue) 暂时不给小节修改meta.json了
				sectionId += 1
				sectionOrder += 1
				docsOrder = 0

				meta := *utils.CreateMeta(info, sectionOrder)
				if Mode == 2 {
					currentSectionsMeta, meta = CreateMetaByCurrentMetas(currentSectionsMeta, path, info, sectionOrder)
				}

				sectionData := models.Section{
					MetaData:   meta,
					ModTime:    info.ModTime(),
					Filepath:   path,
					Locale:     locale,
					ChapterID:  chapterId,
					CategoryID: catId,
					BookID:     bookId,
				}

				sectionDatas = append(sectionDatas, sectionData)

			}
		} else {
			// 暂时不考虑5层的 目前4层够用了
			// 排除一些文件

			if depth >= 4 && utils.PureFileName(info.Name()) != "summary.md" && utils.StringsInside([]string{".md", ".MD"}, info.Name()) {
				docsId += 1
				docsOrder += 1

				meta := utils.CreateMeta(info, docsOrder)

				if Mode == 2 {
					meta, err = utils.ReadMarkdownMeta(path, info, docsOrder)
					if err == nil {
						utils.UpdateMeta(meta, info.Name(), utils.PureFileName(info.Name()), docsOrder, false)
					}
				}

				docsData := models.Document{
					MetaData:   *meta,
					ModTime:    info.ModTime(),
					Locale:     locale,
					CategoryID: catId,
					BookID:     bookId,
					Filepath:   path,
				}

				if depth >= 5 {
					docsData.ChapterID = chapterId
				}
				if depth == 6 {
					docsData.SectionID = sectionId
				}

				docsDatas = append(docsDatas, docsData)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	if Mode <= 1 {
		fmt.Println("正在保存元数据信息到章节...")
		WriteMetadataToLocal(
			currentCatBooks,
			cateLocalMeta,
			currentBookChapters,
			bookLocalMeta,
			summaryMeta,
			catDatas,
		)

		if Mode == 0 {
			fmt.Println("注意: 正在保存元数据信息到文章...")
			go func() {
				for _, docsData := range docsDatas {
					utils.InitMarkdownMeta(docsData)
				}
			}()
		}

	}

	fmt.Println("正在写入数据库...")

	WriteContentToDocsData(&docsDatas)

	fmt.Println("保存数据中...")
	err = WriteIntoDatabase(db,
		interface{}(catDatas),
		interface{}(bookDatas),
		interface{}(chapterDatas),
		interface{}(sectionDatas),
		interface{}(docsDatas))
	if err != nil {
		return err
	}

	fmt.Println("数据库生成成功")
	fmt.Println("用时", time.Since(start))

	if Mode <= 1 {

		global.AppConfig.DBMode = 2
		err := global.WriteConfigToFile("DBMode")
		if err != nil {
			println(err)
			fmt.Println("数据库模式切换失败,请手动切换")
		} else {
			fmt.Println("数据库生成模式已切换为更新模式")
		}

	}

	return nil
}
