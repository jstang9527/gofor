package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/jstang9527/gofor/src/share/utils"
)

// HTTPProxy ...
type HTTPProxy struct {
	target string
}

// Run 返回代理地址
// @param target: http://127.0.0.1:8080
func (p *HTTPProxy) Run() (string, error) {
	// 0.解析后端地址
	backend, err := url.Parse(p.target)
	if err != nil {
		return "", err
	}
	// 1.获取本机随机addr(ip+port)
	ip, err := utils.GetLocalIP()
	if err != nil {
		return "", err
	}
	// 2.创建代理
	reverse := httputil.NewSingleHostReverseProxy(backend)
	server := http.Server{
		Addr:         ip,
		WriteTimeout: 3 * time.Second,
		Handler:      reverse,
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	return ip, nil
}
