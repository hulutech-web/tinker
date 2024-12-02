// Code generated by 'yaegi extract github.com/goravel/framework/database/gorm'. DO NOT EDIT.

package symbol

import (
	"github.com/goravel/framework/contracts/database"
	"github.com/goravel/framework/database/gorm"
	"gorm.io/gorm"
	"reflect"
)

func init() {
	Symbols["github.com/goravel/framework/database/gorm/gorm"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BuildQueryImpl":          reflect.ValueOf(gorm.BuildQueryImpl),
		"DialectorSet":            reflect.ValueOf(&gorm.DialectorSet).Elem(),
		"ErrorMissingWhereClause": reflect.ValueOf(&gorm.ErrorMissingWhereClause).Elem(),
		"FileWithLineNum":         reflect.ValueOf(gorm.FileWithLineNum),
		"GormSet":                 reflect.ValueOf(&gorm.GormSet).Elem(),
		"InitializeGorm":          reflect.ValueOf(gorm.InitializeGorm),
		"InitializeQuery":         reflect.ValueOf(gorm.InitializeQuery),
		"NewDialectorImpl":        reflect.ValueOf(gorm.NewDialectorImpl),
		"NewEvent":                reflect.ValueOf(gorm.NewEvent),
		"NewGormImpl":             reflect.ValueOf(gorm.NewGormImpl),
		"NewLogger":               reflect.ValueOf(gorm.NewLogger),
		"NewMysql1Docker":         reflect.ValueOf(gorm.NewMysql1Docker),
		"NewMysqlDocker":          reflect.ValueOf(gorm.NewMysqlDocker),
		"NewPostgresqlDocker":     reflect.ValueOf(gorm.NewPostgresqlDocker),
		"NewQueryImpl":            reflect.ValueOf(gorm.NewQueryImpl),
		"NewSqliteDocker":         reflect.ValueOf(gorm.NewSqliteDocker),
		"NewSqlserverDocker":      reflect.ValueOf(gorm.NewSqlserverDocker),
		"NewToSql":                reflect.ValueOf(gorm.NewToSql),
		"NewTransaction":          reflect.ValueOf(gorm.NewTransaction),
		"QuerySet":                reflect.ValueOf(&gorm.QuerySet).Elem(),
		"ToCarbonHookFunc":        reflect.ValueOf(gorm.ToCarbonHookFunc),
		"ToDeletedAtHookFunc":     reflect.ValueOf(gorm.ToDeletedAtHookFunc),
		"ToTimeHookFunc":          reflect.ValueOf(gorm.ToTimeHookFunc),

		// type definitions
		"Address":          reflect.ValueOf((*gorm.Address)(nil)),
		"Author":           reflect.ValueOf((*gorm.Author)(nil)),
		"Book":             reflect.ValueOf((*gorm.Book)(nil)),
		"Box":              reflect.ValueOf((*gorm.Box)(nil)),
		"Conditions":       reflect.ValueOf((*gorm.Conditions)(nil)),
		"CursorImpl":       reflect.ValueOf((*gorm.CursorImpl)(nil)),
		"Dialector":        reflect.ValueOf((*gorm.Dialector)(nil)),
		"DialectorImpl":    reflect.ValueOf((*gorm.DialectorImpl)(nil)),
		"Event":            reflect.ValueOf((*gorm.Event)(nil)),
		"GormImpl":         reflect.ValueOf((*gorm.GormImpl)(nil)),
		"Having":           reflect.ValueOf((*gorm.Having)(nil)),
		"House":            reflect.ValueOf((*gorm.House)(nil)),
		"Join":             reflect.ValueOf((*gorm.Join)(nil)),
		"Logger":           reflect.ValueOf((*gorm.Logger)(nil)),
		"MysqlDocker":      reflect.ValueOf((*gorm.MysqlDocker)(nil)),
		"People":           reflect.ValueOf((*gorm.People)(nil)),
		"Person":           reflect.ValueOf((*gorm.Person)(nil)),
		"Phone":            reflect.ValueOf((*gorm.Phone)(nil)),
		"PostgresqlDocker": reflect.ValueOf((*gorm.PostgresqlDocker)(nil)),
		"Product":          reflect.ValueOf((*gorm.Product)(nil)),
		"QueryImpl":        reflect.ValueOf((*gorm.QueryImpl)(nil)),
		"Review":           reflect.ValueOf((*gorm.Review)(nil)),
		"Role":             reflect.ValueOf((*gorm.Role)(nil)),
		"Select":           reflect.ValueOf((*gorm.Select)(nil)),
		"SqliteDocker":     reflect.ValueOf((*gorm.SqliteDocker)(nil)),
		"SqlserverDocker":  reflect.ValueOf((*gorm.SqlserverDocker)(nil)),
		"Table":            reflect.ValueOf((*gorm.Table)(nil)),
		"Tables":           reflect.ValueOf((*gorm.Tables)(nil)),
		"ToSql":            reflect.ValueOf((*gorm.ToSql)(nil)),
		"Transaction":      reflect.ValueOf((*gorm.Transaction)(nil)),
		"User":             reflect.ValueOf((*gorm.User)(nil)),
		"Where":            reflect.ValueOf((*gorm.Where)(nil)),
		"With":             reflect.ValueOf((*gorm.With)(nil)),

		// interface wrapper definitions
		"_Dialector": reflect.ValueOf((*_github_com_goravel_framework_database_gorm_Dialector)(nil)),
	}
}

// _github_com_goravel_framework_database_gorm_Dialector is an interface wrapper for Dialector type
type _github_com_goravel_framework_database_gorm_Dialector struct {
	IValue interface{}
	WMake  func(configs []database.Config) ([]gorm.Dialector, error)
}

func (W _github_com_goravel_framework_database_gorm_Dialector) Make(configs []database.Config) ([]gorm.Dialector, error) {
	return W.WMake(configs)
}