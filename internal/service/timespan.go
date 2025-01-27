package service

func UserPurchaseTimespan(userID int64) error {
	_, err := DBbase.CreateUsertimespan(userID)
	if err != nil {
		return err
	}
	return nil
}

func ResetAllUserTimespanMonthly() error {
	tsList, err := DBbase.GetValidUserTimespanList()
	if err != nil {
		return err
	}
	for _, v := range tsList {
		v.SpendFlow = 0
		DBbase.UpdateUserTimespan(v)
	}
	return nil
}
