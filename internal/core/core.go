package core

import (
	"gostonc/internal/model"
	"gostonc/internal/repo"
)

var (
	Appbase repo.RepoBase

	LoginIDUser map[int64]*model.User = make(map[int64]*model.User)
)

func Init() {
	Appbase = repo.NewRepoBase()

	userList, err := Appbase.GetUserList()
	if err != nil {
		//TODO log fatal err
		return
	}

	for _, u := range userList {
		LoginIDUser[u.ID] = u
	}
}
