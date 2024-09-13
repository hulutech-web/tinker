package tinker

import (
	_ "embed"
)

// 嵌入文件内容
//
//go:embed funcs/db/db.go
var DBFileContent string

// 导出一个函数，供外部调用
func GetDBFileContent() string {
	return DBFileContent
}
