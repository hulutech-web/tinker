package commands

import (
	"fmt"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/support/path"
	"goravel/packages/tinker/view"
	"os"
)

type Tinker struct{}

func NewTinker() *Tinker {
	return &Tinker{}
}

// Signature The name and signature of the console command.
func (receiver *Tinker) Signature() string {
	return "tinker"
}

// Description The console command description.
func (receiver *Tinker) Description() string {
	return "a tinker command tool"
}

// Extend The console command extend.
func (receiver *Tinker) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *Tinker) Handle(ctx console.Context) error {
	// 指定要创建的文件路径和名称
	dirPath := path.Storage("tmp")
	filePath := path.Storage("tmp/test")

	// 先判断目录是否存在，如果不存在则创建目录
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		fmt.Println("创建目录失败:", err)
	}

	// 在已创建或已存在的目录下创建文件
	_, err1 := os.Create(filePath)

	if err1 != nil {
		panic(err1.Error())
	}
	view.Call()
	return nil
}
