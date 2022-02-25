package model

type User struct {
	UserID   int64  `eorm:"user_id"`
	Username string `eorm:"username"`
	Password string `eorm:"password"`
}

type UserResp struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type UserReq struct {
	UserID   string `json:"user_id" binding:"required"`
	Username string `json:"username" binding:"required"`
}
