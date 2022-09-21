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

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return

		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if err := biz.UpdateRestaurant(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}

}
