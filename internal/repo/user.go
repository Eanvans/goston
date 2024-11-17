package repo

import "gostonc/internal/model"

func (r *RepoModule) CreateUser(u *model.User) (*model.User, error) {
	err := r.db.
		Model(&model.User{}).
		Create(u).
		Error
	return u, err
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
