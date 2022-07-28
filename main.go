package main

import (
	"food-delivery-application/component"
	"food-delivery-application/component/uploadprovider"
	"food-delivery-application/middleware"
	"food-delivery-application/modules/restaurant/restauranttransport/ginrestaurant"
	"food-delivery-application/modules/upload/uploadtransport/ginupload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dbConStr := os.Getenv("DBConnectionStr")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	// secretKey := os.Getenv("SYSTEM_SECRET")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(postgres.Open(dbConStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3Provider); err != nil {
		//if err := runService(db); err != nil {
		log.Fatalln(err)

	}
}

func runService(db *gorm.DB, s3Provider uploadprovider.UploadProvider) error {
	//func runService(db *gorm.DB) error {
	r := gin.Default()
	appCtx := component.NewAppContext(db, s3Provider)
	//appCtx := component.NewAppContext(db)

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/upload", ginupload.Upload(appCtx))

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))

		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))

		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))

		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))

	}

	return r.Run()
}
