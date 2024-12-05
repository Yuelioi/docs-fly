package document

import (
	"context"

	"docsfly/internal/svc"
	"docsfly/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDocumentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新文章
func NewUpdateDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDocumentLogic {
	return &UpdateDocumentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDocumentLogic) UpdateDocument(req *types.Document) (resp *types.IDResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
