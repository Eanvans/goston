package service

import (
	"gostonc/internal/app/errcode"
	"gostonc/internal/core"
	"gostonc/internal/model"
)

func UserRegister(username, password string) (*model.User, error) {
	user, err := core.Appbase.CreateUser(model.NewUser(username, password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DoLogin(username string, passowrd string) (*model.User, error) {
	user, err := core.Appbase.GetUserByUsername(username)
	if err != nil {
		return nil, errcode.UnauthorizedAuthNotExist
	}

	if user.Model != nil && user.ID > 0 {
		// 对比密码是否正确
		if model.ValidPassword(user.Password, passowrd, user.Salt) {
			if user.Status == model.USER_STATUS_BANNED {
				return nil, errcode.UserHasBeenBanned
			}
		}

		return nil, errcode.UnauthorizedAuthFailed
	}

	return nil, errcode.UnauthorizedAuthNotExist
}
