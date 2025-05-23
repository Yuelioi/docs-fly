// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	document "docsfly/internal/handler/document"
	"docsfly/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 发布文章
				Method:  http.MethodPost,
				Path:    "/document",
				Handler: document.PushDocumentHandler(serverCtx),
			},
			{
				// 获取页面评论
				Method:  http.MethodGet,
				Path:    "/document/:id",
				Handler: document.DocumentHandler(serverCtx),
			},
			{
				// 更新文章
				Method:  http.MethodPut,
				Path:    "/document/:id",
				Handler: document.UpdateDocumentHandler(serverCtx),
			},
			{
				// 删除文章
				Method:  http.MethodDelete,
				Path:    "/document/:id",
				Handler: document.DeleteDocumentHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}
