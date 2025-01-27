package core

import "gostonc/internal/model"

type ITimespanRepo interface {
	CreateUsertimespan(userID int64) (*model.TimeSpan, error)
	UpdateUserTimespan(ts *model.TimeSpan) error

	GetValidUserTimespanList() ([]*model.TimeSpan, error)
}
