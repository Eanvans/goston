package service

import "gostonc/internal/core"

func UserPurchaseTimespan(userID int64) error {
	_, err := core.Appbase.CreateUsertimespan(userID)
	if err != nil {
		return err
	}
	return nil
}
