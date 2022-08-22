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

func ListUsersLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		paging.Fulfill()

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewListUsersLikeRestaurantBiz(store)

		result, err := biz.ListUsersLikeRestaurant(c.Request.Context(), &filter, &paging, nil)
		if err != nil {
			panic(err)
			return
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))

	}

}
