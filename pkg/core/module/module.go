package module

// Module 模块
type Module interface {
	OnInit() error          // 初始化
	OnDestroy() error       // 销毁
	Run(closeSig chan bool) // 启动
	Name() string           // 名字
}
