package userbiz

import (
	"context"
	"food-delivery-application/common"
	"food-delivery-application/component"
	"food-delivery-application/component/tokenprovider"
	"food-delivery-application/modules/user/usermodel"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBiz struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	customHash    CustomHash
	expiry        int
}

func NewLoginBiz(storeUser LoginStorage, tokenProvider tokenprovider.Provider,
	customHash CustomHash, expiry int) *loginBiz {
	return &loginBiz{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		customHash:    customHash,
		expiry:        expiry,
	}
}

// 1. Find user, email
// 2. Hash pass from input and compare with pass in db
// 3. Provider: issue JWT token for client
// 3.1. Access token and refresh token
// 4. Return token(s)

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.LoginRequest) (*usermodel.LoginResponse, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	hashedPassword := biz.customHash.CustomHash(data.Password + user.Salt)

	if user.Password != hashedPassword {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role.String(),
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewLoginResponse(accessToken, nil)

	return account, nil
}
