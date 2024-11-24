package service

func UserPurchaseTimespan(userID int64) error {
	_, err := DBbase.CreateUsertimespan(userID)
	if err != nil {
		return err
	}
	return nil
}
