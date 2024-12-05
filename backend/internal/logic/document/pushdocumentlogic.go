package document

import (
	"context"

	"docsfly/internal/svc"
	"docsfly/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushDocumentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发布文章
func NewPushDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushDocumentLogic {
	return &PushDocumentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PushDocumentLogic) PushDocument(req *types.Document) (resp *types.IDResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
