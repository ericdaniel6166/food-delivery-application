package restaurantlikemodel

import (
	"fmt"
	"food-delivery-application/common"
	"time"
)

const EntityName = "RestaurantLike"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
}

type LikedCount struct {
	RestaurantId int `gorm:"column:restaurant_id;"`
	LikeCount    int `gorm:"column:count;"`
}

func (Like) TableName() string { return "restaurant_likes" }

func (restaurantLike *Like) GetRestaurantId() int {
	return restaurantLike.RestaurantId
}

func (restaurantLike *Like) GetUserId() int {
	return restaurantLike.UserId
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"),
	)
}

func ErrCannotUnlikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot unlike this restaurant"),
		fmt.Sprintf("ErrCannotUnlikeRestaurant"),
	)
}
