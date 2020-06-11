package imagegenerator

import (
	"io"
	"strings"
	"sync"
	"sync/atomic"

	"errors"
)

type protocolReader struct {
	protocol string
	open     func(url string) (io.ReadCloser, error)
}

// Resource 资源类型
type Resource string

var (
	readerMu        sync.Mutex
	readerProtocols atomic.Value
	// ErrForNotSupportResource 不支持的资源类型错误
	ErrForNotSupportResource error = errors.New("not support resource protocol")
)

// RegisterReader 注册一个资源读取器.
// Protocol 读取器支持的协议，如 http:// 、https://、local:// 等
// Open 打开资源并返回一个读取器
func RegisterReader(protocol string, open func(url string) (io.ReadCloser, error)) {
	readerMu.Lock()
	protocols, _ := readerProtocols.Load().([]protocolReader)
	readerProtocols.Store(append(protocols, protocolReader{protocol, open}))
	readerMu.Unlock()
}

// Sniff determines the protocol of url.
func sniffResource(url string) protocolReader {
	protocols, _ := readerProtocols.Load().([]protocolReader)
	for _, f := range protocols {
		if strings.HasPrefix(url, f.protocol) {
			return f
		}
	}
	return protocolReader{}
}

// Open 打开资源
func (r Resource) Open() (io.ReadCloser, error) {
	url := string(r)
	p := sniffResource(url)
	if p.protocol == "" {
		return nil, ErrForNotSupportResource
	}
	return p.open(url)
}
