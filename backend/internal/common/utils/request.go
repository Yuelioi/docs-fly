package utils

import (
	"net"
	"net/http"
	"strings"
)

// getIPFromRequest 从请求中提取客户端的 IP 地址
func GetIPFromRequest(r *http.Request) string {
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		// X-Forwarded-For 可能包含多个 IP 地址，取第一个
		return strings.Split(ip, ",")[0]
	}

	// 如果没有 X-Forwarded-For 头，从 X-Real-Ip 头中获取 IP 地址
	if ip := r.Header.Get("X-Real-Ip"); ip != "" {
		return ip
	}

	// 如果没有 X-Real-Ip 头，从 RemoteAddr 中获取 IP 地址
	if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return ip
	}

	return ""
}
