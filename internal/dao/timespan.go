package dao

import (
	"gostonc/internal/model"
	"time"
)

func (r *RepoModule) CreateUsertimespan(userID int64) (*model.TimeSpan, error) {
	u := &model.TimeSpan{}
	//BY Default
	err := r.db.
		Model(&model.TimeSpan{}).
		Create(&model.TimeSpan{
			UserID:     userID,
			TotalFlow:  1024 * 1024 * 1024 * 1024,
			SpendFlow:  0,
			ExpireDate: time.Now().Unix(),
		}).Error

	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *RepoModule) UpdateUserTimespan(ts *model.TimeSpan) error {
	err := r.db.
		Model(ts).
		Where("id = ? AND is_del = ?", ts.ID, 0).
		Save(ts).
		Error
	return err
}
