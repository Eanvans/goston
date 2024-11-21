package repo

import (
	"gostonc/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type RepoModule struct {
	db *gorm.DB
}

type RepoBase interface {
	IUserRepo
	ITimespanRepo
	IAuthenticate
}

type IUserRepo interface {
	CreateUser(u *model.User) (*model.User, error)
	UpdateUser(u *model.User) error
	GetUserByUsername(username string) (*model.User, error)

	GetUserList() ([]*model.User, error)
}

type IAuthenticate interface {
	Authenticate(username, password string) (bool, *model.User)
}

type ITimespanRepo interface {
	CreateUsertimespan(userID int64) (*model.TimeSpan, error)
	UpdateUserTimespan(ts *model.TimeSpan) error
}

func NewRepoBase() RepoBase {
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
