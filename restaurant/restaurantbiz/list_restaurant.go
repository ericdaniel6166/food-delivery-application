package restaurantbiz

import (
	"context"
	"food-delivery-application/common"
	"food-delivery-application/restaurant/restaurantmodel"
)

type ListRestaurantStore interface {
	ListDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		order *common.Order,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	order *common.Order,
) ([]restaurantmodel.Restaurant, error) {

	if err := order.Validate(); err != nil {
		return nil, err

	}

	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging, order)

	return result, err
}
