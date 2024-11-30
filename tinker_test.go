package tinker

import (
	"github.com/traefik/yaegi/interp"
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

	// 检查Tinker实例的interpreter字段是否为interp.New()的返回值
	if tinkerInstance.interpreter != interp.New(interp.Options{}) {
		t.Errorf("NewTinker() returned a Tinker instance with a different interpreter")
	}
}
