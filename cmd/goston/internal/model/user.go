package model

// 用户类型
type UserType int

const (
	USER_TYPE_NORMAL UserType = iota
	USER_TYPE_VIP
	USER_TYPE_SVIP
	USER_TYPE_ADMIN = 99
)

type User struct {
	*Model
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`

	Status int    `json:"status"`
	Avatar string `json:"avatar"`

	Type UserType `json:"type"` // 用户类型
}
