package proxy

import (
	"fmt"
	"log"

	"github.com/jstang9527/gofor/src/share/server/tcp"
	"github.com/jstang9527/gofor/src/share/utils"
)

// TCPProxy ...
type TCPProxy struct {
	target string
}

// Run 返回代理地址
// @param target: 127.0.0.1:6379
func (p *TCPProxy) Run() (string, error) {
	// 1.获取本机随机addr(ip+port)
	ip, err := utils.GetLocalIP()
	if err != nil {
		return "", err
	}
	// 2.创建代理
	// reverse := httputil.NewSingleHostReverseProxy(backend)
	// server := http.Server{
	// 	Addr:         ip,
	// 	WriteTimeout: 3 * time.Second,
	// 	Handler:      reverse,
	// }
	ss1 := NewTcpReverseProxy(p.target)
	tcpServ := tcp.TcpServer{Addr: ip, Handler: ss1}
	go func() {
		fmt.Println("Starting tcp_proxy at " + ip)
		log.Fatal(tcpServ.ListenAndServe())
	}()
	return ip, nil
}
