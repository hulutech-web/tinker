package tinker

import (
	_ "embed"
)

// 嵌入文件内容
//
//go:embed funcs/db/db.go
var DBFileContent string

//go:embed funcs/query/query.go
var QueryFileContent string

// 导出一个函数，供外部调用
func GetDBFileContent() string {
	return DBFileContent
}
func GetQueryFileContent() string {
	return QueryFileContent
}
