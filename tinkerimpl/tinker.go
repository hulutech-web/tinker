package tinkerimpl

import (
	"fmt"
	"github.com/hulutech-web/goravel-tinker/contracts"
)

var tinker contracts.Tinker

type TinkerImpl struct {
	tinker contracts.Tinker
}

func NewTinkerImpl() *TinkerImpl {
	return &TinkerImpl{
		tinker: tinker,
	}
}

func (t *TinkerImpl) Call() {
	t.tinker.Call()
}
func (t *TinkerImpl) GetQueryFileContent() string {
	fmt.Println("Tinker")
	return ""
}

func (t *TinkerImpl) GetDBFileContent() string {
	fmt.Println("Tinker")
	return ""
}
