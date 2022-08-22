package restaurantmodel

import (
	"errors"
	"food-delivery-application/common"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	UserId          int                `json:"-" gorm:"column:user_id;"`
	User            *common.SimpleUser `json:"user,omitempty" gorm:"preload:false;"`
	Address         string             `json:"address" gorm:"column:address;"`
	Logo            *common.Image      `json:"logo,omitempty" gorm:"column:logo;"`
	Cover           *common.Images     `json:"cover,omitempty" gorm:"column:cover;"`
	LikedCount      int                `json:"liked_count,omitempty" gorm:"column:liked_count;"` // computed field
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	common.SQLModel `json:",inline"`
	Name            *string        `json:"name" gorm:"column:name;"`
	Address         *string        `json:"address" gorm:"column:address;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	UserId          int            `json:"-" gorm:"column:user_id;"`
	Address         string         `json:"address" gorm:"column:address;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (restaurantCreate RestaurantCreate) Validate() error {
	restaurantCreate.Name = strings.TrimSpace(restaurantCreate.Name)
	if len(restaurantCreate.Name) == 0 {
		return errors.New("restaurant name can't be blank")
	}
	return nil
}

func (restaurant *Restaurant) Mask(isAdminOrOwner bool) {
	restaurant.GenUID(common.DbTypeRestaurant)
	if user := restaurant.User; user != nil {
		user.Mask(isAdminOrOwner)
	}
}

func (restaurantCreate *RestaurantCreate) Mask(isAdminOrOwner bool) {
	restaurantCreate.GenUID(common.DbTypeRestaurant)
}
