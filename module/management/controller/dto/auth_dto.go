package dto

type LoginReq struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterReq struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
