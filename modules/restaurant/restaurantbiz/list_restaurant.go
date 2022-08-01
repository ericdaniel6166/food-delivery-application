package restaurantbiz

import (
	"context"
	"food-delivery-application/common"
	"food-delivery-application/modules/restaurant/restaurantmodel"
)

type ListRestaurantRepo interface {
	ListRestaurant(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		order *common.Order,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{repo: repo}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	order *common.Order,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.repo.ListRestaurant(ctx, filter, paging, order)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
