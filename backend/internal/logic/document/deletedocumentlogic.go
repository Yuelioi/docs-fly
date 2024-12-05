package document

import (
	"context"

	"docsfly/internal/svc"
	"docsfly/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDocumentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除文章
func NewDeleteDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDocumentLogic {
	return &DeleteDocumentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDocumentLogic) DeleteDocument(req *types.Document) (resp *types.IDResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
