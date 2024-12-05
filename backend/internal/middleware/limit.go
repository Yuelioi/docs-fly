package middleware

import (
	"docsfly/internal/common/utils"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Limit() func(next http.HandlerFunc) http.HandlerFunc {
	type accessRecord struct {
		count    int
		lastTime time.Time
	}

	accessRecords := make(map[string]accessRecord)
	var mu sync.Mutex

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ip := utils.GetIPFromRequest(r)

			mu.Lock()
			record, exists := accessRecords[ip]
			if !exists {
				record = accessRecord{count: 0, lastTime: time.Now()}
			}

			// 检查是否超过 1 分钟
			if time.Since(record.lastTime) > time.Minute {
				record.count = 0
			}

			// 检查访问次数
			if record.count >= 360 {
				httpx.ErrorCtx(r.Context(), w, errors.New("请求太多啦, 请稍后重试"))
				mu.Unlock()
				return
			}

			// 更新访问记录
			record.count++
			record.lastTime = time.Now()
			accessRecords[ip] = record
			mu.Unlock()

			// 继续处理请求
			next.ServeHTTP(w, r)
		}
	}
}
