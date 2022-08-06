package userbiz

import (
	"context"
	"food-delivery-application/common"
	"food-delivery-application/component/salt"
	"food-delivery-application/modules/user/usermodel"
)

type RegisterStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type CustomHash interface {
	CustomHash(data string) string
}

type registerBiz struct {
	registerStore RegisterStore
	customHash    CustomHash
}

func NewRegisterBiz(registerStore RegisterStore, customHash CustomHash) *registerBiz {
	return &registerBiz{
		registerStore: registerStore,
		customHash:    customHash,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, err := biz.registerStore.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err == common.RecordNotFound {
		genSalt := salt.GenSalt(50)

		data.Password = biz.customHash.CustomHash(data.Password + genSalt)
		data.Salt = genSalt
		data.Role = common.User.String()
		data.Status = 1

		if err := biz.registerStore.CreateUser(ctx, data); err != nil {
			return common.ErrCannotCreateEntity(usermodel.EntityName, err)
		}

		return nil
	} else if user != nil {
		return usermodel.ErrEmailExisted
	} else {
		return err
	}

}
