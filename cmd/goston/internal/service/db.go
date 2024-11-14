package service

import (
	"gostonc/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	SqliteDB *gorm.DB
)

func InitDatabase() {
	// // 打开或创建 SQLite 数据库
	// db, err := sql.Open(dbVersion, dbPath)
	// if err != nil {
	// 	log.Fatalf("Failed to open database: %v", err)
	// }

	// 初始化数据库连接
	db, err := gorm.Open(sqlite.Open("goston.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模式
	db.AutoMigrate(&model.User{})

	SqliteDB = db
}
