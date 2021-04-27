package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/jstang9527/gofor/src/share/config"
	"github.com/jstang9527/gofor/src/share/path"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

// var consulReg = consul.NewRegistry(registry.Addrs(config.ConsulAddr))

type Gateway struct {
	c client.Client
}

// 处理具体的rpc请求
func (g *Gateway) handleJSONRPC(w http.ResponseWriter, r *http.Request) {
	// 处理请求路径, 得到具体服务和方法、将url转换为service和method
	service, method := path.PathToReceiver(config.Namespace, r.URL.Path)
	log.Println("service:"+service, "method:"+method)

	// 读取请求体
	br, _ := ioutil.ReadAll(r.Body)
	// 封装request
	request := json.RawMessage(br)
	// 调用服务
	var response json.RawMessage
	req := g.c.NewRequest(service, method, &request, client.WithContentType("application/json"))
	err := g.c.Call(path.RequestToContext(r), req, &response)

	out := ""
	if err != nil {
		// 出错,把错误信息返回给客户端
		out = err.Error()
	} else {
		b, _ := response.MarshalJSON()
		out = string(b)
	}
	w.Header().Set("Content-Length", strconv.Itoa(len([]byte(out))))
	if _, err = w.Write([]byte(out)); err != nil {
		log.Println(err)
	}
}

func main() {
	// consulService := micro.NewService(micro.Registry(consulReg))
	consulService := micro.NewService()
	consulService.Init()
	gateway := &Gateway{c: consulService.Client()}

	mux := http.NewServeMux()
	mux.HandleFunc("/", gateway.handleJSONRPC)
	log.Println("Start server at ::8888")
	log.Fatal(http.ListenAndServe(":8888", mux))
}
