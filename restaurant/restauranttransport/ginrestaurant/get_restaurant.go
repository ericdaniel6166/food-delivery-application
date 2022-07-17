package ginrestaurant

import (
	"food-delivery-application/common"
	"food-delivery-application/component"
	"food-delivery-application/restaurant/restaurantbiz"
	"food-delivery-application/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			//c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			panic(common.ErrInvalidRequest(err))
			return

		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		result, err := biz.GetRestaurant(c.Request.Context(), id)
		if err != nil {
			panic(err)
			return
		}
		// test

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))

	}

}
