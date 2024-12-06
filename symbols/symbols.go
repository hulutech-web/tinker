//go:generate yaegi extract github.com/goravel/framework/database/orm
//go:generate yaegi extract github.com/goravel/framework/facades
//go:generate yaegi extract github.com/mitchellh/mapstructure
//go:generate yaegi extract gorm.io/driver/mysql
//go:generate yaegi extract gorm.io/gorm
//go:generate yaegi extract goravel/packages/tinker/funcs/auth
//go:generate yaegi extract goravel/packages/tinker/funcs/cache
//go:generate yaegi extract goravel/packages/tinker/funcs/command
//go:generate yaegi extract goravel/packages/tinker/funcs/db
//go:generate yaegi extract goravel/packages/tinker/funcs/query
//go:generate yaegi extract goravel/app/models

package symbols

import "reflect"

var Symbols = map[string]map[string]reflect.Value{}
