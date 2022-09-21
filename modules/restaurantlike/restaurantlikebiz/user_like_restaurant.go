package restaurantlikebiz

import (
	"context"
	"food-delivery-application/common"
	"food-delivery-application/modules/restaurantlike/restaurantlikemodel"
	"food-delivery-application/pubsub"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
}

//type IncreaseLikeCountStore interface {
//	IncreaseLikeCount(ctx context.Context, id int) error
//}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	//incStore IncreaseLikeCountStore
	pubsub pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore,
	//incStore IncreaseLikeCountStore,
	pubsub pubsub.Pubsub,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store,
		//incStore: incStore,
		pubsub: pubsub,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.RestaurantLike,
) error {
	// Find, if present, return "already like"

	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	//// side effect

	//job := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//})
	//
	//_ = asyncjob.NewGroup(true, job).Run(ctx)

	// New solution: Use pub/sub

	biz.pubsub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))

	return nil
}
