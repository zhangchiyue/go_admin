package app

import (
	"adx-admin/pkg/core/module"
	"adx-admin/pkg/database/redshift"
	"adx-admin/pkg/microlog"
	"adx-admin/pkg/server"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
)

// 节点全局状态
const (
	StateNone = iota // 未开始或已停止
	StateInit        // 正在初始化中
	StateRun         // 正在运行中
	StateStop        // 正在停止中
)

// mod 模块
type mod struct {
	mi       module.Module
	closeSig chan bool
	wg       sync.WaitGroup
}

// NewApp 创建App
func NewApp(redshift *redshift.Redshift, srv *server.Server, logger microlog.Logger) *App {
	app := &App{
		closeSig: make(chan os.Signal, 1),
		state:    StateNone,
		mods:     make([]*mod, 0, 2),
		log:      logger,
	}
	// todo 怎么优化掉重复的工作
	app.mods = append(app.mods,
		&mod{
			mi:       redshift,
			closeSig: make(chan bool, 1),
		}, &mod{
			mi:       srv,
			closeSig: make(chan bool, 1),
		})
	return app
}

// App 中的 modules 在初始化(通过 Start 或 Run) 之后不能变更
// App 有两种启停方式:
//  1. Start -> Stop: 手动启动和停止app，比较干净，通常用于测试代码
//  2. Run -> Terminate: 基于Start/Stop封装，自动监听OS Signal或通过Terminate来终止，通常用于真正的节点启动流程
type App struct {
	mods                    []*mod
	state                   int32
	closeSig                chan os.Signal
	wg                      sync.WaitGroup
	beforeModulesRunHook    func() error
	afterModulesDestroyHook func()
	log                     microlog.Logger
}

func (app *App) SetBeforeModulesRunHook(h func() error) {
	app.beforeModulesRunHook = h
}

func (app *App) SetAfterModulesDestroyHook(h func()) {
	app.afterModulesDestroyHook = h
}

// SetState 设置状态
func (app *App) setState(s int32) {
	atomic.StoreInt32(&app.state, s)
}

// GetState 获取状态
func (app *App) GetState() int32 {
	return atomic.LoadInt32(&app.state)
}

// Start 非阻塞启动app，需要在当前goroutine调用Stop来停止app
func (app *App) Start() {
	// 单个app不能启动两次
	if app.GetState() != StateNone {
		app.log.Fatal("app mods cannot start twice")
	}

	app.wg.Add(1)
	// 注册module 并增加开关
	app.log.Info("app is starting...")
	// register

	app.setState(StateInit)

	// 模块初始化
	for _, m := range app.mods {
		mi := m.mi
		if err := mi.OnInit(); err != nil {
			app.log.Fatalf("module %v init error %v", reflect.TypeOf(mi), err)
		}
	}

	if app.beforeModulesRunHook != nil {
		if err := app.beforeModulesRunHook(); err != nil {
			app.log.Fatalf("failed to exec afterModulesInitHook: %v", err)
		}
	}

	// 模块启动
	for _, m := range app.mods {
		m.wg.Add(1)
		go run(m)
	}
	app.setState(StateRun)
	app.log.Info("app is started")
}

// Stop 停止App
func (app *App) Stop() {
	if app.GetState() == StateStop {
		return
	}
	app.setState(StateStop)
	app.log.Infof("app is stopping...")
	// 先进后出
	for i := len(app.mods) - 1; i >= 0; i-- {
		m := app.mods[i]
		close(m.closeSig)
		m.wg.Wait()
		destroy(m)
	}
	if app.afterModulesDestroyHook != nil {
		app.afterModulesDestroyHook()
	}

	app.mods = nil
	app.wg.Done()
	app.setState(StateNone)
	app.log.Infof("app is stopped")
}

// run 启动所有模块
func run(m *mod) {
	defer printPanicStack()
	defer m.wg.Done()
	m.mi.Run(m.closeSig)
}

// destroy 销毁模块
func destroy(m *mod) error {
	defer printPanicStack()
	return m.mi.OnDestroy()
}

// Run 阻塞启动app，在监测到SIGINT SIGTERM信号时自动终止App
// 也可在任意goroutine调用Terminate来终止
func (app *App) Run() {
	app.Start()
	// 信号监听 优雅退出
	for {
		signal.Notify(app.closeSig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
		sig := <-app.closeSig
		if sig == syscall.SIGHUP {
			continue
		}
		break
	}
	app.Stop()
}

// Terminate 用于模拟信号，终止Run，并等待app停止完成
// goroutine safe
func (app *App) Terminate() {
	if app.GetState() != StateRun {
		return
	}
	app.closeSig <- syscall.SIGQUIT
}

func printPanicStack() {
	if r := recover(); r != nil {
		buf := make([]byte, 2048)
		l := runtime.Stack(buf, false)
		logrus.Errorf("%v: %s", r, buf[:l])
		fmt.Printf("%v: %s\n", r, buf[:l])
	}
}
