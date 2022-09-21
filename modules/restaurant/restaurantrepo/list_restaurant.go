package restaurantrepo

import (
	"context"
	"food-delivery-application/common"
	"food-delivery-application/modules/restaurant/restaurantmodel"
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

type LikeStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantRepo struct {
	store     ListRestaurantStore
	likeStore LikeStore
}

func NewListRestaurantRepo(store ListRestaurantStore, likeStore LikeStore) *listRestaurantRepo {
	return &listRestaurantRepo{
		store:     store,
		likeStore: likeStore,
	}
}

func (repo *listRestaurantRepo) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	order *common.Order,
) ([]restaurantmodel.Restaurant, error) {
	if err := order.Validate(); err != nil {
		return nil, err
	}
	result, err := repo.store.ListDataByCondition(ctx, nil, filter, paging, order, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}
	return result, nil
}
