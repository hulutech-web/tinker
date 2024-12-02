package query

import (
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
)

var result map[string]interface{}
var results []map[string]interface{}

func Query() orm.Query {
	return facades.Orm().Query()
}

func _clearResults() {
	results = []map[string]interface{}{}
}
func _clearResult() {
	result = map[string]interface{}{}
}

func Clear() {
	_clearResults()
	_clearResult()
}
