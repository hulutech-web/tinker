package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("tinker", map[string]any{
		"name": "Tinker",
	})
}
