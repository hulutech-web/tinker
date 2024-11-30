package facade

import (
	tinker "github.com/hulutech-web/goravel-tinker"
)

// TinkerFacade 是一个全局单例
var TinkerFacade *Tinker

type Tinker struct {
	instance *tinker.Tinker
}

func NewTinker() *Tinker {
	return &Tinker{
		instance: tinker.NewTinker(),
	}
}

func (t *Tinker) Call() {
	// 调用实际的 tinker 实现
	// 直接调用 view.Call()，与根目录 tinker.go 中的实现保持一致
	t.instance.Call()
}

// 确保能访问到嵌入的文件内容
func (t *Tinker) GetDBFileContent() string {
	return t.instance.GetDBFileContent()
}

func (t *Tinker) GetQueryFileContent() string {
	return t.instance.GetQueryFileContent()
}

// 初始化函数
func init() {
	TinkerFacade = NewTinker()
}
