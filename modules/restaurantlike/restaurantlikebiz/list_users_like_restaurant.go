package restaurantlikebiz

import (
	"context"
	"food-delivery-application/common"
	"food-delivery-application/modules/restaurantlike/restaurantlikemodel"
)

type ListUsersLikeRestaurantStore interface {
	GetUsersLikeRestaurant(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		order *common.Order,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUsersLikeRestaurantBiz struct {
	store ListUsersLikeRestaurantStore
}

func NewListUsersLikeRestaurantBiz(store ListUsersLikeRestaurantStore) *listUsersLikeRestaurantBiz {
	return &listUsersLikeRestaurantBiz{store: store}
}

func (biz *listUsersLikeRestaurantBiz) ListUsersLikeRestaurant(
	ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	order *common.Order,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUsersLikeRestaurant(ctx, nil, filter, paging, order)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return users, nil
}
