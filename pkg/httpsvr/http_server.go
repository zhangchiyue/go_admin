package httpsvr

import (
	"context"
	"net"
	"net/http"
)

// HTTPServer HTTP服务器
type HTTPServer struct {
	Addr    string
	Handler http.Handler
	server  *http.Server
}

// Start 启动HttpServer
func (hs *HTTPServer) Start() error {
	ln, err := net.Listen("tcp", hs.Addr)
	if err != nil {
		return err
	}

	hs.server = &http.Server{
		Addr:    hs.Addr,
		Handler: hs.Handler,
	}
	go func() {
		_ = hs.server.Serve(ln)
	}()
	return nil
}

// Shutdown 关闭HTTP服务器
func (hs *HTTPServer) Shutdown(ctx context.Context) error {
	if hs.server != nil {
		return hs.server.Shutdown(ctx)
	} else {
		return nil
	}
}
