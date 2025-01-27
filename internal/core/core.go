package core

type RepoBase interface {
	IUserRepo
	ITimespanRepo
	IAuthenticate
}
