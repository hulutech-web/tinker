package query

import (
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 定义一个测试用的结构体，并指定表名为 users
type TestUser struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"type:varchar(100)"`
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete
}

// TableName 指定表名为 users
func (TestUser) TableName() string {
	return "users"
}

func setupTestDB() *gorm.DB {
	// 使用 MySQL 数据库进行测试
	dsn := "root:Dazhouquxian2012.@tcp(127.0.0.1:3306)/goravel-workflow?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect to database")
	}

	// 自动迁移
	db.AutoMigrate(&TestUser{})
	return db
}

// 清空 users 表
func clearUsersTable(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE users") // 清空 users 表
}

func TestCreateWithMap(t *testing.T) {
	db := setupTestDB()
	gormIns = db // 使用测试数据库

	// 清空 users 表
	clearUsersTable(db)

	// 使用 map 创建用户，手动设置 CreatedAt 和 UpdatedAt
	userData := map[string]interface{}{
		"Name":      "jinzhu",
		"Age":       18,
		"CreatedAt": time.Now(),
		"UpdatedAt": time.Now(),
	}

	result := db.Model(&TestUser{}).Create(userData)

	// 检查插入结果
	if result.Error != nil {
		t.Fatalf("Error inserting user: %v", result.Error)
	}

	// 查询最后插入的用户
	var user TestUser
	if err := db.Last(&user).Error; err != nil {
		t.Fatalf("Failed to find the last inserted user: %v", err)
	}

	// 验证 ID 是否匹配
	if user.ID == 0 {
		t.Fatalf("Expected created user to have a non-zero ID, got %d", user.ID)
	}

	// 验证其他字段
	if user.Name != userData["Name"] || user.Age != userData["Age"] {
		t.Errorf("Expected user %v, got %v", userData, user)
	}
}

func TestFind(t *testing.T) {
	db := setupTestDB()
	gormIns = db // 使用测试数据库

	// 清空 users 表
	clearUsersTable(db)

	// 插入测试数据
	users := []TestUser{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 28},
	}
	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			t.Fatalf("Failed to create test user: %v", err)
		}
	}

	// 测试 Find 函数
	var results []TestUser
	Find(&results) // 直接传入切片

	if len(results) != len(users) {
		t.Errorf("Expected %d results, got %d", len(users), len(results))
	}
}
