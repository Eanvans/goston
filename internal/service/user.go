package service

import (
	"gostonc/internal/app/errcode"
	"gostonc/internal/model"
)

func UserRegister(username, password string) (*model.User, error) {
	user, err := DBbase.CreateUser(model.NewUser(username, password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DoLogin(username string, passowrd string) (*model.User, error) {
	user, err := DBbase.GetUserByUsername(username)
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

func PurchaseTimespan(userID int64) error {
	timespan, err := DBbase.CreateUsertimespan(userID)
	if err != nil {
		return err
	}

	user, err := DBbase.GetUserByID(userID)

	if u, ok := loginUnameUser[user.Username]; ok {
		u.TimeSpan = *timespan
	}
	return nil
}

func AuthUserByUsernamePassword(username, password string) (bool, *model.User) {
	if user, ok := loginUnameUser[username]; ok {
		return ok, user
	}

	ok, user := DBbase.Authenticate(username, password)
	if ok {
		loginUnameUser[username] = user
		return ok, user
	}

	return false, nil
}

func GetCacheUserByUname(username string) (*model.User, bool) {
	if user, ok := loginUnameUser[username]; ok {
		return user, ok
	}
	return nil, false
}
