package server

import (
	"adx-admin/pkg/configer"
	"adx-admin/pkg/httpsvr"
	"adx-admin/pkg/microlog"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func NewServer(config *configer.Config, engine *gin.Engine, log microlog.Logger) *Server {
	addr := ":" + strconv.Itoa(config.HttpServerConfig.HttpPort)
	return &Server{
		HTTPPort: config.HttpServerConfig.HttpPort,
		httpServer: &httpsvr.HTTPServer{
			Addr:    addr,
			Handler: engine,
		},
		log: log,
	}
}

type Server struct {
	HTTPPort   int
	httpServer *httpsvr.HTTPServer
	log        microlog.Logger
}

func (m *Server) OnInit() error {
	return nil
}

// OnDestroy 销毁模块
func (m *Server) OnDestroy() error {
	return m.stop()
}

// Run 启动
func (m *Server) Run(closeSig chan bool) {
	err := m.httpServer.Start()
	if err != nil {
		m.httpServer = nil
		m.log.Panicf("start http server failed: %v", err)
	} else {
		m.log.Infof("start http server successful on port: %v", m.HTTPPort)
	}

}

func (m *Server) stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := m.httpServer.Shutdown(ctx); err != nil {
		m.log.Errorf("shutdown http server error: %v", err)
		return err
	} else {
		m.log.Info("shutdown http server successfully")
	}
	return nil
}

// Name 获取模块的名字
func (m *Server) Name() string {
	return "server"
}
