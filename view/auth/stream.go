package auth

import "github.com/goravel/framework/facades"

func Store(key string, value string) {
	facades.Cache().Put(key, value, 0)
}

func Get(key string) any {
	return facades.Cache().Get(key)
}

func RetString(str string) string {
	return str
}
