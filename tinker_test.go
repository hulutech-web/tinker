package tinker

import (
	"testing"
)

func TestNewTinker(t *testing.T) {
	// 创建一个新的Tinker实例
	tinkerInstance := NewTinker()

	// 检查Tinker实例是否为nil
	if tinkerInstance == nil {
		t.Errorf("NewTinker() returned nil")
	}

	// 检查Tinker实例的interpreter字段是否为nil
	if tinkerInstance.interpreter == nil {
		t.Errorf("NewTinker() returned a Tinker instance with a nil interpreter")
	}
}

func TestCall(t *testing.T) {
	// 创建一个新的Tinker实例
	tinkerInstance := NewTinker()
	//调用Call方法
	tinkerInstance.Call()
}
