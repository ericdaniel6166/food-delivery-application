package ginrestaurant

import (
	"food-delivery-application/common"
	"food-delivery-application/component"
	"food-delivery-application/modules/restaurant/restaurantbiz"
	"food-delivery-application/modules/restaurant/restaurantmodel"
	"food-delivery-application/modules/restaurant/restaurantrepo"
	"food-delivery-application/modules/restaurant/restaurantstorage"
	"food-delivery-application/modules/restaurantlike/restaurantlikestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		paging.Fulfill()

		var order common.Order

		if err := c.ShouldBind(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		order.Fulfill()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		likeStore := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		repo := restaurantrepo.NewListRestaurantRepo(store, likeStore)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging, &order)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))

	}

}
