// Code generated by 'yaegi extract tinker/funcs/command'. DO NOT EDIT.

//go:build go1.22
// +build go1.22

package symbols

import (
	"reflect"
	"tinker/funcs/command"
)

func init() {
	Symbols["tinker/funcs/command/command"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AppkeyGenerate": reflect.ValueOf(command.AppkeyGenerate),
		"Call":           reflect.ValueOf(command.Call),
		"JwtGenerate":    reflect.ValueOf(command.JwtGenerate),
	}
}
