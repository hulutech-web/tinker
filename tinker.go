package tinker

import (
	"github.com/hulutech-web/goravel-tinker/internal/executor"
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

func (T *Tinker) Call() error {
	return executor.Call()
}
