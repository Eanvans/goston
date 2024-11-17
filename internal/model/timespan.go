package model

type TimeSpan struct {
	*Model
	UserID int64 `gorm:"column:user_id;uniqueIndex" json:"user_id"`

	TotalFlow int64 `json:"total_flow"` //总流量
	SpendFlow int64 `json:"spend_flow"` //消耗流量

	ExpireDate int64 `json:"expire_date"` //到期时间
}
