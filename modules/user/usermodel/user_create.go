package usermodel

import "food-delivery-application/common"

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string          `json:"email" gorm:"column:email;"`
	Password        string          `json:"password" gorm:"column:password;"`
	LastName        string          `json:"last_name" gorm:"column:last_name;"`
	FirstName       string          `json:"first_name" gorm:"column:first_name;"`
	Role            *common.AppRole `json:"-" gorm:"column:role;type:string"`
	Salt            string          `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image   `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}
