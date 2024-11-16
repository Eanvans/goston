package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gofrs/uuid"
)

// 用户类型
type UserType int

const (
	USER_TYPE_NORMAL UserType = iota
	USER_TYPE_VIP
	USER_TYPE_SVIP
	USER_TYPE_ADMIN = 99
)

type UserStatus int

const (
	USER_STATUS_NORMAL UserStatus = iota
)

type User struct {
	*Model
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`

	Status UserStatus `json:"status"`
	Avatar string     `json:"avatar"`

	Type UserType `json:"type"` // 用户类型
}

func NewUser(username, passowrd string) *User {
	password, salt := encryptPasswordAndSalt(passowrd)

	return &User{
		Nickname: username,
		Username: username,
		Password: password,
		Salt:     salt,
		Status:   USER_STATUS_NORMAL,
	}
}

func encryptPasswordAndSalt(password string) (string, string) {
	salt := uuid.Must(uuid.NewV4()).String()[:8]
	password = encodeMD5(encodeMD5(password) + salt)

	return password, salt
}

func encodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
