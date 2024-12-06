package tinker

import (
	"fmt"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type Tinker struct {
	interpreter *interp.Interpreter
}

func NewTinker() *Tinker {
	return &Tinker{
		interpreter: interp.New(interp.Options{}),
	}
}

func (t *Tinker) StartUp() {
	i := t.interpreter

	i.Use(stdlib.Symbols)
	addstr := Add(1, 2)
	res, err := i.Eval(addstr)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Interface().(int))
	if err != nil {
		panic(err)
	}
}

func Add(a, b int) string {
	addstr :=
		`package foo
	func Add(a, b int) int { return a + b }`
	return addstr
}
