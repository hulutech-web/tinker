package query

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

var dest interface{}
var result interface{}

type Gorm struct {
	gormIns *gorm.DB
}

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

func DB() *gorm.DB {
	return BootMS()
}

// users := []*User{ {Name: "Jinzhu", Age: 18, Birthday: time.Now()},}
// eg: result := db.Create(users) // pass a slice to insert multiple row

func Model(value interface{}) (tx *gorm.DB) {
	dest = value
	return DB().Model(result)
}

func Create(value interface{}) interface{} {
	Clear()
	dest = value
	DB().Create(&dest)
	return dest
}

func Select(query interface{}, args ...interface{}) interface{} {
	return DB().Select(query, args)
}

func Take(dest interface{}, conds ...interface{}) interface{} {
	return DB().Take(dest, conds)
}
func First(dest interface{}, conds ...interface{}) interface{} {
	return DB().First(dest, conds)
}
func Last(dest interface{}, conds ...interface{}) interface{} {
	return DB().Last(dest, conds)
}

// users := []models.User{}
//
//eg:query.Find(&users)
func Find(dest interface{}, conds ...interface{}) interface{} {
	DB().Find(dest, conds)
	return dest
}

func Where(query interface{}, args ...interface{}) interface{} {
	return DB().Where(query, args)
}
func Or(query interface{}, args ...interface{}) interface{} {
	return DB().Or(query, args)
}
func Order(value interface{}) interface{} {
	return DB().Order(value)
}
func Limit(limit int) interface{} {
	return DB().Limit(limit)
}
func Group(name string) interface{} {
	return DB().Group(name)
}
func Distinct(cols ...interface{}) interface{} {
	return DB().Distinct(cols...)
}
func Joins(table string, args ...interface{}) interface{} {
	return DB().Joins(table, args...)
}
func Scan(dest interface{}) (tx *gorm.DB) {
	return DB().Scan(dest)
}

func _clearResult() {
	result = map[string]interface{}{}
}

func Clear() {
	_clearResult()
}
