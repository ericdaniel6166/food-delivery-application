package subscriber

import (
	"context"
	"food-delivery-application/component"
	"food-delivery-application/modules/restaurant/restaurantstorage"
	"food-delivery-application/pubsub"
)

type HasRestaurantId interface {
	GetRestaurantId() int
}

//func IncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserLikeRestaurant)
//
//	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
//
//	go func() {
//		defer common.AppRecover()
//		for {
//			msg := <-c
//			likeData := msg.Data().(HasRestaurantId)
//			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
//		}
//	}()
//}

//// I wish I could do something like that
//func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext) func(ctx context.Context, message *pubsub.Message) error {
//	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
//
//	return func(ctx context.Context, message *pubsub.Message) error {
//		likeData := message.Data().(HasRestaurantId)
//		return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
//	}
//}

func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
