package cache

import (
	"fmt"
	"github.com/goravel/framework/facades"
)

func Store(key string, value string) {
	err := facades.Cache().Put(key, value, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		panic(err)
	}
}

func Get(key string) any {
	return facades.Cache().Get(key)
}

func RetString(str string) string {
	return str
}
