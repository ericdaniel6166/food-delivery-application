package restaurantbiz

import (
	"context"
	"food-delivery-application/common"
	restaurantmodel2 "food-delivery-application/modules/restaurant/restaurantmodel"
)

type ListRestaurantStore interface {
	ListDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel2.Filter,
		paging *common.Paging,
		order *common.Order,
		moreKeys ...string,
	) ([]restaurantmodel2.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel2.Filter,
	paging *common.Paging,
	order *common.Order,
) ([]restaurantmodel2.Restaurant, error) {

	if err := order.Validate(); err != nil {
		return nil, err

	}

	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging, order)

	return result, err
}
