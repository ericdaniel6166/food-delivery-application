package usermodel

import "food-delivery-application/common"

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string         `json:"email" gorm:"column:email;"`
	Password        string         `json:"-" gorm:"column:password;"`
	Salt            string         `json:"-" gorm:"column:salt;"`
	LastName        string         `json:"last_name" gorm:"column:last_name;"`
	FirstName       string         `json:"first_name" gorm:"column:first_name;"`
	Phone           string         `json:"phone" gorm:"column:phone;"`
	Role            common.AppRole `json:"role" gorm:"column:role;"`
	Avatar          *common.Image  `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
	IsOnline        bool           `json:"is_online" gorm:"-"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role.String()
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}
