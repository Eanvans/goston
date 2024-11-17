package core

import "gostonc/internal/repo"

var (
	Appbase repo.RepoBase
)

func Init() {
	Appbase = repo.NewRepoBase()
}
