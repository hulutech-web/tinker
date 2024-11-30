package tinker

import "github.com/traefik/yaegi/interp"

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

// 嵌入 db.go 文件
//
//go:embed funcs/db/db.go
var DBFileContent string

// 嵌入 query.go 文件
//
//go:embed funcs/query/query.go
var QueryFileContent string

// 导出一个函数，供外部调用
func GetDBFileContent() string {
	return DBFileContent
}

func GetQueryFileContent() string {
	return QueryFileContent
}
