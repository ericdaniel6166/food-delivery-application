package ginuser

import (
	"food-delivery-application/common"
	"food-delivery-application/component"
	"food-delivery-application/component/customhash"
	"food-delivery-application/modules/user/userbiz"
	"food-delivery-application/modules/user/usermodel"
	"food-delivery-application/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := customhash.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
