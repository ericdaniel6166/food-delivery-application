package main

import (
	"food-delivery-application/component"
	"food-delivery-application/component/uploadprovider"
	"food-delivery-application/middleware"
	"food-delivery-application/modules/restaurant/restauranttransport/ginrestaurant"
	ginrestaurantlike "food-delivery-application/modules/restaurantlike/restaurantliketransport/ginrestaurantlike"
	"food-delivery-application/modules/upload/uploadtransport/ginupload"
	"food-delivery-application/modules/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	dbConStr := os.Getenv("DBConnectionStr")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SystemSecretKey")
	version := os.Getenv("Version")
	jwtExpirationInSeconds, _ := strconv.Atoi(os.Getenv("JwtExpirationInSeconds"))

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(postgres.Open(dbConStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	appCtx := component.NewAppContext(db, s3Provider, secretKey, version, jwtExpirationInSeconds)

	db = db.Debug()

	if err := runService(appCtx); err != nil {
		log.Fatalln(err)
	}
}

func runService(appCtx component.AppContext) error {
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	v := r.Group(appCtx.Version())

	v.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v.POST("/upload", ginupload.Upload(appCtx))

	auth := v.Group("/auth")
	{
		auth.POST("/register", ginuser.Register(appCtx))
		auth.POST("/login", ginuser.Login(appCtx))
	}

	user := v.Group("/user", middleware.RequiredAuth(appCtx))
	{
		user.GET("/profile", ginuser.GetProfile(appCtx))
	}

	restaurants := v.Group("/restaurant", middleware.RequiredAuth(appCtx))
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))

		restaurants.GET("/:id/liked-users", ginrestaurantlike.ListUsersLikeRestaurant(appCtx))
		restaurants.POST("/:id/like", ginrestaurantlike.UserLikeRestaurant(appCtx))
		restaurants.DELETE("/:id/unlike", ginrestaurantlike.UserUnlikeRestaurant(appCtx))

	}

	return r.Run()
}
