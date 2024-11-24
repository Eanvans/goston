package core

import "gostonc/internal/model"

type IUserRepo interface {
	CreateUser(u *model.User) (*model.User, error)
	UpdateUser(u *model.User) error
	GetUserByUsername(username string) (*model.User, error)
	GetUserByID(id int64) (*model.User, error)

	GetUserList() ([]*model.User, error)
}
