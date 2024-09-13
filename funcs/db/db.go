package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	once sync.Once
)

// 申明一个MYSQL连接GormIns
var gormIns *gorm.DB

func BootMS() *gorm.DB {
	once.Do(func() {
		ins, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       "root:Dazhouquxian2012.@tcp(127.0.0.1:3306)/goravel-workflow?charset=utf8&parseTime=True&loc=Local", // DSN data source name
			DefaultStringSize:         256,                                                                                                 // string 类型字段的默认长度
			DisableDatetimePrecision:  true,                                                                                                // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,                                                                                                // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,                                                                                                // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false,                                                                                               // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic(err)
		}
		gormIns = ins
	})
	return gormIns
}

var result map[string]interface{}
var results []map[string]interface{}

type Gorm struct {
	gormIns *gorm.DB
}

func DB() *gorm.DB {
	return BootMS()
}

func Take(table string) map[string]interface{} {
	Clear()
	DB().Table(table).Take(&result)
	return result
}

// eg:db.Find("depts")
func Find(table string) ([]map[string]interface{}, int64) {
	Clear()
	tx := DB().Table(table).Find(&results)
	return results, tx.RowsAffected
}

// eg:db.WhereMap("users",map[string]interface{}{"name":"市场部-经理"})
func WhereMap(table string, expr map[string]interface{}) ([]map[string]interface{}, int64) {
	Clear()
	tx := DB().Table(table).Where(expr).Find(&results)
	return results, tx.RowsAffected
}

//eg:db.WhereNot("users",map[string]interface{}{"name":"市场部-经理"})
func WhereNot(table string, expr map[string]interface{}) ([]map[string]interface{}, int64) {
	Clear()
	tx := DB().Table(table).Not(expr).Limit(10).Find(&results)

	return results, tx.RowsAffected
}

//eg:db.WhereNotEqual("users","name","市场部-经理")
func WhereNotEqual(table string, column string, condition string) ([]map[string]interface{}, int64) {
	Clear()
	tx := DB().Table(table).Where(column+" <> ?", condition).Find(&results)
	return results, tx.RowsAffected
}

// eg:db.WhereIn("users","name", []string{"市场部-经理", "市场部-主管"})
func WhereIn(table string, column string, conditions []string) ([]map[string]interface{}, int64) {
	Clear()
	tx := DB().Table(table).Where(column+" IN ?", conditions).Find(&results)
	return results, tx.RowsAffected
}

//eg:db.WhereLike("users","name", "市场部")
func WhereLike(table string, column string, condition string) ([]map[string]interface{}, int64) {
	Clear()
	tx := DB().Table(table).Where(column+" like ?", "%"+condition+"%").Find(&results)
	return results, tx.RowsAffected
}

//eg:db.WhereBetween("users","created_at", `2024-08-22 00:00:00`, `2024-08-23 23:59:59`)
func WhereBetween(table string, column string, start interface{}, end interface{}) ([]map[string]interface{}, int64) {
	Clear()
	tx := DB().Table(table).Where(column+" BETWEEN ? AND ?", start.(string), end.(string)).Find(&results)
	return results, tx.RowsAffected
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
