package dao

import (
	"gostonc/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type RepoModule struct {
	db *gorm.DB
}

func NewRepoBase() *RepoModule {
	db, err := gorm.Open(sqlite.Open("goston.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//ensureDBDelete(db)

	// 自动迁移模式
	db.AutoMigrate(
		&model.User{},
		&model.TimeSpan{},
		&model.Order{},
	)

	rm := &RepoModule{
		db: db,
	}

	return rm
}

func ensureDBDelete(db *gorm.DB) {
	// 获取 Migrator
	migrator := db.Migrator()
	// 清除现有表（可选）
	if migrator.HasTable(&model.User{}) {
		migrator.DropTable(&model.User{})
	}
	if migrator.HasTable(&model.TimeSpan{}) {
		migrator.DropTable(&model.TimeSpan{})
	}
	if migrator.HasTable(&model.Order{}) {
		migrator.DropTable(&model.Order{})
	}
}
