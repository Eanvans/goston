package repo

import (
	"gostonc/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type RepoModule struct {
	db *gorm.DB
}

func (r *RepoModule) CreateUser(u *model.User) (*model.User, error) {
	err := r.db.
		Model(&model.User{}).
		Create(u).
		Error
	return u, err
}

type RepoBase interface {
	IUserRepo
}

type IUserRepo interface {
	CreateUser(u *model.User) (*model.User, error)
}

func NewRepoBase() RepoBase {
	db, err := gorm.Open(sqlite.Open("goston.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模式
	db.AutoMigrate(&model.User{})

	rm := &RepoModule{
		db,
	}

	return rm
}
