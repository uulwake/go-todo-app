package handler

type UserRegisterRequest struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=20"`
}
type UserLoginRequest struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=20"`
}

type UserRegisterLoginResponseData struct {
	UserID int64 `json:"user_id"`
	UserName string `json:"user_name"`
	Token string `json:"jwt_token"`
}

type UserRegisterLoginResponse struct {
	Data UserRegisterLoginResponseData `json:"data"`
}
