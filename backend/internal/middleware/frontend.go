package middleware

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/zeromicro/go-zero/rest"
)

// 开启前后端一起部署

const basename = "/web" // 虚拟路由根路径

// //go:embed public
var assets embed.FS

type NotFoundHandler struct {
	fs         http.FileSystem
	fileServer http.Handler
}

func (n NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := strings.TrimPrefix(path.Clean(r.URL.Path), basename)
	if len(filePath) == 0 {
		filePath = basename
	}

	file, err := n.fs.Open(filePath)
	switch {
	case err == nil:
		n.fileServer.ServeHTTP(w, r)
		_ = file.Close()
		return
	case os.IsNotExist(err):
		r.URL.Path = "/"
		n.fileServer.ServeHTTP(w, r)
		return
	default:
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
}

func Frontend() func(*rest.Server) {
	sub, _ := fs.Sub(assets, "public")
	fs := http.FS(sub)
	fileServer := http.FileServer(fs)

	return rest.WithNotFoundHandler(&NotFoundHandler{
		fs:         fs,
		fileServer: fileServer,
	})
}
