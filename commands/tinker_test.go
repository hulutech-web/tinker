package commands

import (
	"testing"
)

func TestTinker_Handle(t *testing.T) {
	// 创建 Tinker 命令实例
	tinkerCmd := NewTinker()

	// 调用 Handle 方法
	err := tinkerCmd.Handle(nil)
	if err != nil {
		t.Errorf("Handle returned unexpected error: %v", err)
	}
}
