package repo

import (
	"gostonc/internal/model"
)

func (r *RepoModule) CreateUser(u *model.User) (*model.User, error) {
	err := r.db.
		Model(&model.User{}).
		Create(u).
		Error
	return u, err
}

func (r *RepoModule) UpdateUser(u *model.User) error {
	err := r.db.
		Model(u).
		Where("id = ? AND is_del = ?", u.ID, 0).
		Save(u).
		Error
	return err
}

func (r *RepoModule) GetUserByUsername(username string) (*model.User, error) {
	u := &model.User{}

	err := r.db.
		Where("username = ?", username).
		First(u).
		Error

	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *RepoModule) GetUserList() ([]*model.User, error) {
	var users []*model.User

	err := r.db.
		Preload("TimeSpan").
		Find(&users).
		Error

	return users, err
}

func (r *RepoModule) Authenticate(username, password string) (bool, *model.User) {
	user, err := r.GetUserByUsername(username)
	if err != nil {
		return false, nil
	}

	if user.Model != nil && user.ID > 0 {
		// 对比密码是否正确
		if model.ValidPassword(user.Password, password, user.Salt) {
			if user.Status == model.USER_STATUS_BANNED {
				return false, nil
			}
			return true, user
		}
		return false, nil
	}

	return false, nil
}
