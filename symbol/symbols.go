//go:generate go install github.com/traefik/yaegi/cmd/yaegi@v0.15.0
//go:generate yaegi extract github.com/goravel/framework/facades
//go:generate yaegi extract github.com/goravel/framework/database
//go:generate yaegi extract gorm.io/driver/mysql
//go:generate yaegi extract gorm.io/gorm
//go:generate yaegi extract gorm.io/gorm/logger
//go:generate yaegi extract github.com/spf13/viper

package symbol

import "reflect"

var Symbols = map[string]map[string]reflect.Value{}
