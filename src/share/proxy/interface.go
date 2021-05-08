package proxy

type ProxyType int

const (
	ProxyTypeTCP = iota + 1
	ProxyTypeHTTP
)

// Proxy ...
type Proxy interface {
	Run() (string, error)
}

// NewProxy ...
func NewProxy(proxyType ProxyType, target string) Proxy {
	switch proxyType {
	case ProxyTypeTCP:
		return &TCPProxy{target: target}
	case ProxyTypeHTTP:
		return &HTTPProxy{target: target}
	default: // 默认TCP
		return &TCPProxy{target: target}
	}
}
