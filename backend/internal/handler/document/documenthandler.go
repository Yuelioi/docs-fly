package document

import (
	"net/http"

	"docsfly/internal/logic/document"
	"docsfly/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取页面评论
func DocumentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := document.NewDocumentLogic(r.Context(), svcCtx)
		resp, err := l.Document()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
