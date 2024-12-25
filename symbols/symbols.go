//go:generate go install github.com/traefik/yaegi/cmd/yaegi@v0.15.0
//go:generate yaegi extract github.com/goravel/framework/facades
//go:generate yaegi extract gorm.io/driver/mysql
//go:generate yaegi extract gorm.io/gorm
//go:generate yaegi extract gorm.io/gorm/logger
//go:generate yaegi extract goravel-tinker/funcs/cache
//go:generate yaegi extract goravel-tinker/funcs/command
//go:generate yaegi extract goravel-tinker/funcs/db
//go:generate yaegi extract goravel-tinker/funcs/query
//go:generate yaegi extract github.com/goravel/framework/database/orm

package symbols

import "reflect"

var Symbols = map[string]map[string]reflect.Value{}
