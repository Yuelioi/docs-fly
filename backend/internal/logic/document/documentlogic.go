package document

import (
	"context"

	"docsfly/internal/svc"
	"docsfly/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DocumentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取页面评论
func NewDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DocumentLogic {
	return &DocumentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DocumentLogic) Document() (resp *types.Document, err error) {
	// todo: add your logic here and delete this line

	return
}
