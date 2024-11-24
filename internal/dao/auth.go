package dao

import "gostonc/internal/model"

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
