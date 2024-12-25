package db

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 定义一个测试用的结构体
type TestModel struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100)"`
}

func setupTestDB() *gorm.DB {
	// 使用 SQLite 内存数据库进行测试
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect to database")
	}

	// 自动迁移
	db.AutoMigrate(&TestModel{})
	return db
}

func TestStoreAndGet(t *testing.T) {
	db := setupTestDB()
	// 使用 GORM 的 DB 函数
	gormIns = db

	// 测试数据存储
	testData := TestModel{Name: "Test Name"}
	if err := db.Create(&testData).Error; err != nil {
		t.Fatalf("Failed to create test data: %v", err)
	}

	// 测试数据获取
	var result TestModel
	if err := db.First(&result, testData.ID).Error; err != nil {
		t.Fatalf("Failed to get test data: %v", err)
	}

	if result.Name != testData.Name {
		t.Errorf("Expected %s, got %s", testData.Name, result.Name)
	}
}

func TestFind(t *testing.T) {
	db := setupTestDB()
	gormIns = db

	// 插入测试数据
	testData := []TestModel{
		{Name: "Test Name 1"},
		{Name: "Test Name 2"},
	}
	for _, data := range testData {
		if err := db.Create(&data).Error; err != nil {
			t.Fatalf("Failed to create test data: %v", err)
		}
	}

	// 测试 Find 函数
	_, count := Find("users") // 假设表名为 test_models
	if count != int64(len(testData)) {
		t.Errorf("Expected %d results, got %d", len(testData), count)
	}
}
