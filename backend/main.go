package main

import (
	"context"
	"errors"
	"net/http"

	"docsfly/internal/common/biz"
	"docsfly/internal/common/constants"
	"docsfly/internal/handler"
	"docsfly/internal/middleware"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	_ "docsfly/internal/common/db"
)

func main() {
	logc.Info(context.Background(), "启动服务中")

	server := rest.MustNewServer(
		constants.ConfInst.RestConf,

		// 注册前端端口
		// middleware.Frontend(),

		// 验证失败处理
		rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
			httpx.ErrorCtx(r.Context(), w, biz.AuthError)
		}),
	)

	// 中间件
	server.Use(middleware.CORS())
	server.Use(middleware.Limit())

	defer server.Stop()

	handler.RegisterHandlers(server, constants.SVCInst)

	// 统一错误处理
	httpx.SetErrorHandler(func(err error) (int, any) {
		var e *biz.Error
		switch {
		case errors.As(err, &e):
			return http.StatusOK, biz.Fail(e)
		default:
			return http.StatusOK, biz.Result{
				Code: -500,
				Msg:  err.Error(),
				Data: nil,
			}
		}

	})

	// 统一成功处理
	httpx.SetOkHandler(func(ctx context.Context, data any) any {
		return biz.Success(data)
	})

	logc.Info(context.Background(), "服务启动完毕")
	server.Start()

}
