package api

type RegisterUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserPurchaseReq struct {
	UserID int64 `json:"user_id" binding:"required"`
	//GoodID int64 `json:"good_id" binding:"required"`
}
