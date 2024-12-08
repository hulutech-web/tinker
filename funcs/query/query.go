package query

import (
	"fmt"
	"github.com/goravel/framework/facades"
	"github.com/mitchellh/mapstructure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"sync"
)

var (
	once sync.Once
)
var queryResult map[string]interface{}
var queryResults []map[string]interface{}

// 申明一个MYSQL连接GormIns
var QueryIns *gorm.DB

type MyQuery struct {
	query *gorm.DB
}

func Query() *MyQuery {
	return &MyQuery{
		query: BootMS(),
	}
}

var ptr interface{}

func BootMS() *gorm.DB {
	once.Do(func() {
		DB_HOST := facades.Config().Env("DB_HOST")
		DB_PORT := facades.Config().Env("DB_PORT")
		DB_DATABASE := facades.Config().Env("DB_DATABASE")
		DB_USERNAME := facades.Config().Env("DB_USERNAME")
		DB_PASSWORD := facades.Config().Env("DB_PASSWORD")
		//root:Dazhouquxian2012.@tcp(127.0.0.1:3306)/goravel-workflow?charset=utf8&parseTime=True&loc=Local
		ins, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE), // DSN data source name
			DefaultStringSize:         256,                                                                                                                               // string 类型字段的默认长度
			DisableDatetimePrecision:  true,                                                                                                                              // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,                                                                                                                              // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,                                                                                                                              // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false,                                                                                                                             // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		QueryIns = ins
	})
	return QueryIns
}

func (q *MyQuery) Model(value interface{}) *MyQuery {
	ptr = value
	QueryIns.Model(ptr)
	return q
}

func (q *MyQuery) Where(query interface{}, args ...interface{}) *MyQuery {
	QueryIns = QueryIns.Where(query, args)
	return q
}

func (q *MyQuery) FindOne(dest interface{}, conds ...interface{}) (map[string]interface{}, int64) {
	tx := QueryIns.Find(dest, conds)
	queryResult = StructToMap(dest)
	return queryResult, tx.RowsAffected
}

func (q *MyQuery) FindAll(destPrt interface{}, conds ...interface{}) ([]map[string]interface{}, int64) {
	tx := QueryIns.Find(destPrt, conds)

	value := reflect.ValueOf(destPrt)
	// 如果是指针类型，获取其指向的元素的反射值
	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}
	// 判断是否为切片类型
	if value.Kind() != reflect.Slice {
		return nil, 0
	}
	// 遍历切片元素
	for i := 0; i < value.Len(); i++ {
		element := value.Index(i).Interface()
		// 将单个结构体元素转换为map[string]interface{}
		elementMap := StructToMap(element)
		queryResults = append(queryResults, elementMap)
	}
	return queryResults, tx.RowsAffected
}

func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	err := mapstructure.Decode(obj, &result)
	if err != nil {
		fmt.Println("StructToMap转换出错:", err)
	}
	return result
}
