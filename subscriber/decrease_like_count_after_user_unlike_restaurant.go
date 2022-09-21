package subscriber

import (
	"context"
	"food-delivery-application/component"
	"food-delivery-application/modules/restaurant/restaurantstorage"
	"food-delivery-application/pubsub"
)

func RunDecreaseLikeCountAfterUserUnlikeRestaurant(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user unlikes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
