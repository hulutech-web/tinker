package tinker

import (
	"github.com/hulutech-web/goravel-tinker/view"
	"github.com/traefik/yaegi/interp"
)

import (
	_ "embed"
)

type Tinker struct {
	interpreter *interp.Interpreter
}

func NewTinker() *Tinker {
	return &Tinker{
		interpreter: interp.New(interp.Options{}),
	}
}

func (T *Tinker) Call() {
	view.Call()
}

// 嵌入 db.go 文件
//
//go:embed funcs/db/db.go
var DBFileContent string

// 嵌入 query.go 文件
//
//go:embed funcs/query/query.go
var QueryFileContent string

// 导出一个函数，供外部调用
func (T *Tinker) GetDBFileContent() string {
	return DBFileContent
}

func (T *Tinker) GetQueryFileContent() string {
	return QueryFileContent
}
