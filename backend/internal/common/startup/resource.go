package startup

/*
初始化一些文件
*/

import (
	"docsfly/internal/common/constants"
	"log"
	"os"
	"path/filepath"
)

func initDst(inReviewDir, inReviewMeta, defaultContent string) error {
	// 创建文件夹
	if _, err := os.Stat(inReviewDir); os.IsNotExist(err) {
		err = os.MkdirAll(inReviewDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// 初始化元数据
	if _, err := os.Stat(inReviewMeta); os.IsNotExist(err) {
		f, err := os.OpenFile(inReviewMeta, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.Write([]byte(
			defaultContent,
		))
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	// 初始化审核区
	inReviewDocumentsDir := filepath.Join(constants.ConfInst.Resource.Pending)
	inReviewDocumentsMeta := filepath.Join(inReviewDocumentsDir, constants.ConfInst.Resource.MetaFile)
	inReviewDocumentContent := `title: 待审核
cid: InReview
order: 1
collections:
`
	err := initDst(inReviewDocumentsDir, inReviewDocumentsMeta, inReviewDocumentContent)
	if err != nil {
		log.Fatal("创建审核文件失败")
	}

	// 初始化评论区
	commentsDir := constants.ConfInst.Resource.Comments
	commentsMeta := filepath.Join(commentsDir, constants.ConfInst.Resource.MetaFile)
	commentsContent := `comments:    
    - nickname: "system"
      content: "系统评论测试"
      date: "2024-11-19"
`
	err = initDst(commentsDir, commentsMeta, commentsContent)
	if err != nil {
		log.Fatal("创建评论文件失败")
	}

	// 初始化公告区
	announcesDir := constants.ConfInst.Resource.Comments
	announcesMeta := filepath.Join(announcesDir, constants.ConfInst.Resource.MetaFile)
	announcesContent := `announces:
  - title: 小贴士
    content: 单击收藏可以添加到"我的收藏"
    date: 2024/11/16`
	err = initDst(announcesDir, announcesMeta, announcesContent)
	if err != nil {
		log.Fatal("创建公告失败")
	}
}
