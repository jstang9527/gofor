package proxy

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

func test() {
	server, err := net.Listen("tcp", ":1080")
	if err != nil {
		fmt.Printf("Listen failed: %v\n", err)
		return
	}

	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Printf("Accept failed: %v", err)
			continue
		}
		go process(client)
	}
}

func process(client net.Conn) {
	defer client.Close()
	if err := Socks5Auth(client); err != nil {
		log.Println("Socks5Auth failed, err: ", err)
		return
	}
	target, err := Socks5Connect(client)
	if err != nil {
		log.Println("Socks5Connect failed, err: ", err)
		return
	}
	Socks5Forward(client, target)
}

// Socks5Auth
// https://blog.csdn.net/michael__li/article/details/53941358
// VER: Socks的版本，Socks5默认为0x05，其固定长度为1个字节
// NMETHODS: 第三个字段METHODS的长度，其固定长度为1个字节
// METHODS:	客户端支持的验证方式，可以有多种, 1-255个字节
func Socks5Auth(client net.Conn) error {
	// 1.读客户端消息体
	buf := make([]byte, 256)

	// 1.1 读取VER和NMTHODS
	n, err := io.ReadFull(client, buf[:2])
	if n != 2 {
		return errors.New("invalid auth header, err: " + err.Error())
	}
	ver, nmetheds := int(buf[0]), int(buf[1])
	if ver != 5 {
		return errors.New("invalid version, err: " + err.Error())
	}

	// 1.2 读取METHEDD列表
	n, err = io.ReadFull(client, buf[:nmetheds])
	if n != nmetheds {
		return errors.New("invalid methods, err: " + err.Error())
	}

	// 2.响应客户端消息体(返回协议版本和auth方法即可)
	n, err = client.Write([]byte{0x05, 0x00})
	if n != 2 || err != nil {
		return errors.New("invalid res-write, err: " + err.Error())
	}
	return nil
}

// Socks5Connect 建立连接; 通知Socks服务端, 客户端需要访问哪个远程服务器
func Socks5Connect(client net.Conn) (net.Conn, error) {
	// 1. 读取客户端消息体
	// @VER	0x05，协议版本, 其固定长度为1个字节
	// @CMD 连接方式, 0x01=CONNECT, 0x02=BIND, 0x03=UDP ASSOCIATE, 其固定长度为1个字节
	// @RSV 保留字段, 其固定长度为1个字节
	// @ATYP 地址类型, 0x01=IPv4，0x03=域名，0x04=IPv6, 其固定长度为1个字节
	// @DST.ADDR 目标地址, 根据ATYP进行解析，值长度不定。
	// @DST.PORT 目标端口, 网络字节序（network octec order）其固定长度为2个字节
	buf := make([]byte, 256)

	// 1.1 读取VER、CMD、RSV、ATYP
	n, err := io.ReadFull(client, buf[:4])
	if n != 4 {
		return nil, errors.New("invalid connect header, err: " + err.Error())
	}
	ver, cmd, _, atyp := buf[0], buf[1], buf[2], buf[3]
	if ver != 5 || cmd != 1 {
		// 暂时不支持bind、udp associate
		return nil, errors.New("invalid ver/cmd")
	}
	// 1.2 依据AYTP读取DST.ADDR
	// 0x01：4个字节，对应 IPv4 地址
	// 0x02：首位字节表示域名长度，其后跟着域名字节
	// 0x03：16个字节，对应 IPv6 地址
	addr := ""
	switch atyp {
	case 1:
		n, err = io.ReadFull(client, buf[:4])
		if n != 4 {
			return nil, errors.New("invalid IPv4, err: " + err.Error())
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	case 3:
		n, err = io.ReadFull(client, buf[:1])
		if n != 1 {
			return nil, errors.New("invalid domainLen, err: " + err.Error())
		}
		domainLen := int(buf[0])
		n, err := io.ReadFull(client, buf[:domainLen])
		if n != domainLen {
			return nil, errors.New("invalid domain, err: " + err.Error())
		}
		addr = string(buf[:n])
	case 4:
		// 暂不支持ipv6
		return nil, errors.New("IPv6: no supported yet")
	default:
		return nil, errors.New("invalid ATYP")
	}

	// 1.4读取DST.PORT
	n, err = io.ReadFull(client, buf[:2])
	if n != 2 {
		return nil, errors.New("invalid DST.PORT, err: " + err.Error())
	}

	port := binary.BigEndian.Uint16(buf[:2])
	fmt.Println(addr, port)
	// 2. 建立目标连接
	dstAddrPort := fmt.Sprintf("%s:%d", "192.168.27.129", 8080)
	// dstAddrPort := fmt.Sprintf("%s:%d", addr, port)
	dst, err := net.Dial("tcp", dstAddrPort)
	if err != nil {
		return nil, errors.New("failed dial dst, err: " + err.Error())
	}

	// 3. 响应客户端，隧道连接就绪
	// VER 协议版本, 其固定长度为1个字节
	// REP 状态码, 0x00=成功, 0x01=未知错误，…… , 其固定长度为1个字节
	// RSV 保留字段, 其固定长度为1个字节
	// ATYP 地址类型, 其固定长度为1个字节
	// BND.ADDR 服务器和DST创建连接用的地址
	// BND.PORT 服务器和DST创建连接用的端口

	// 3.1 本机隧道ip和端口(由于传回去，客户端也用不上, 就像net.DialTCP()的laddr)
	// laip, laport, err := net.SplitHostPort(dst.LocalAddr().String())
	// fmt.Println(laip, laport, err)

	// 3.2 开始响应
	n, err = client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
	if err != nil {
		dst.Close()
		return nil, errors.New("failed res-write, err: " + err.Error())
	}
	return dst, nil
}

func Socks5Forward(client, target net.Conn) {
	forward := func(src, dest net.Conn) {
		defer src.Close()
		defer dest.Close()
		io.Copy(src, dest)
	}
	go forward(client, target)
	forward(target, client)
}
