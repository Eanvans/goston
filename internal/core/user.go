package core

import "gostonc/internal/model"

type IAuthenticate interface {
	Authenticate(username, password string) (bool, *model.User)
}

type IUserRepo interface {
	CreateUser(u *model.User) (*model.User, error)
	UpdateUser(u *model.User) error
	GetUserByUsername(username string) (*model.User, error)
	GetUserByID(id int64) (*model.User, error)

	GetUserList() ([]*model.User, error)
}
