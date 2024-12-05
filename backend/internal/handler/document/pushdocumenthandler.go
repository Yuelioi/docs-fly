package document

import (
	"net/http"

	"docsfly/internal/logic/document"
	"docsfly/internal/svc"
	"docsfly/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 发布文章
func PushDocumentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Document
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := document.NewPushDocumentLogic(r.Context(), svcCtx)
		resp, err := l.PushDocument(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
