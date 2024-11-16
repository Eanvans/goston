package service

import (
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
