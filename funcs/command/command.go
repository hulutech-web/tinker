package command

import "github.com/goravel/framework/facades"

func JwtGenerate() {
	facades.Artisan().Call("jwt:generate")
}

func AppkeyGenerate() {
	facades.Artisan().Call("key:generate")
}

func Call(command string) {
	facades.Artisan().Call(command)
}
