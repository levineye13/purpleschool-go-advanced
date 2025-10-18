package auth

type TLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type TRegisterRequest struct {
	Name string `json:"name" validate:"required"`
	*TLoginRequest
}

type TLoginResponse struct {
	Token string `json:"token"`
}
