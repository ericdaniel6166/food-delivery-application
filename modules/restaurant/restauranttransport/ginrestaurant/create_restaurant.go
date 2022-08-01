package ginrestaurant

import (
	"food-delivery-application/common"
	"food-delivery-application/component"
	"food-delivery-application/modules/restaurant/restaurantbiz"
	"food-delivery-application/modules/restaurant/restaurantmodel"
	"food-delivery-application/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}

}
