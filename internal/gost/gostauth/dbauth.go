package gostauth

import (
	"gostonc/internal/model"
	"gostonc/internal/service"
	"sync"
	"time"
)

type DBAuthenticator struct {
	User *model.User //已Auth的用户 需要定期清理？

	// 保留
	period  time.Duration
	stopped chan struct{}
	mux     sync.RWMutex
}

func NewDBAuthenticator() *DBAuthenticator {
	return &DBAuthenticator{
		stopped: make(chan struct{}),
	}
}

func (au *DBAuthenticator) Authenticate(username, password string) bool {
	if au == nil {
		return false
	}
	au.mux.RLock()
	defer au.mux.RUnlock()

	ok, user := service.AuthUserByUsernamePassword(username, password)
	if !ok {
		return false
	}
	au.User = user
	return true
}

func (au *DBAuthenticator) GetCurrentUser() *model.User {
	return au.User
}
