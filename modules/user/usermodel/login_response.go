package usermodel

import (
	"food-delivery-application/component/tokenprovider"
)

type LoginResponse struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewLoginResponse(accessToken, refreshToken *tokenprovider.Token) *LoginResponse {
	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
