package ginrestaurant

import (
	"food-delivery-application/common"
	"food-delivery-application/component"
	"food-delivery-application/modules/restaurantlike/restaurantlikebiz"
	"food-delivery-application/modules/restaurantlike/restaurantlikemodel"
	"food-delivery-application/modules/restaurantlike/restaurantlikestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.RestaurantLike{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		//incStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		//biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store, incStore)
		biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store, appCtx.GetPubsub())

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
