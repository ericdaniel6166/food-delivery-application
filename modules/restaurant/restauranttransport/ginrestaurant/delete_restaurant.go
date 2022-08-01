package ginrestaurant

import (
	"food-delivery-application/common"
	"food-delivery-application/component"
	"food-delivery-application/modules/restaurant/restaurantbiz"
	"food-delivery-application/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return

		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}

}
