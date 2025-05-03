package models

type AuthLoginRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type AuthLoginResponse struct {
	AccessToken string `json:"access_token"`
}