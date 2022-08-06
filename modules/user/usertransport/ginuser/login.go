package ginuser

import (
	"food-delivery-application/common"
	"food-delivery-application/component"
	"food-delivery-application/component/customhash"
	"food-delivery-application/component/tokenprovider/jwt"
	"food-delivery-application/modules/user/userbiz"
	"food-delivery-application/modules/user/usermodel"
	"food-delivery-application/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest usermodel.LoginRequest

		if err := c.ShouldBind(&loginRequest); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := customhash.NewMd5Hash()

		biz := userbiz.NewLoginBiz(store, tokenProvider, md5, appCtx.JwtExpirationInSeconds())
		loginResponse, err := biz.Login(c.Request.Context(), &loginRequest)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(loginResponse))
	}
}
