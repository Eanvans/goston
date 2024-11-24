package core

import (
	"gostonc/internal/model"
)

type RepoBase interface {
	IUserRepo
	ITimespanRepo
	IAuthenticate
}

type IAuthenticate interface {
	Authenticate(username, password string) (bool, *model.User)
}

type ITimespanRepo interface {
	CreateUsertimespan(userID int64) (*model.TimeSpan, error)
	UpdateUserTimespan(ts *model.TimeSpan) error
}
