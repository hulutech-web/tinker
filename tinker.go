package tinker

import (
	"github.com/traefik/yaegi/interp"
)

type Tinker struct {
	interpreter *interp.Interpreter
}

func NewTinker() *Tinker {
	return &Tinker{
		interpreter: interp.New(interp.Options{}),
	}
}
