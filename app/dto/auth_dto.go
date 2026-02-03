package dto

type LoginAuthDto struct {
	Username string `json:"username" validate:"required,min=6,max=100"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
