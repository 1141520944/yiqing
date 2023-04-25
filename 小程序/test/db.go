package test

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	User      User   `gorm:"foreignkey:u_id;association_foreignkey:id"`
	UID       uint   `gorm:"column:u_id" json:"-"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	Phone     string `gorm:"column:phone" json:"phone"`
	Address   string `gorm:"column:address" json:"address"`
}

func (UserInfo) TableName() string {
	return "t_user_info"
}

type User struct {
	gorm.Model
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"-"`
}

func (User) TableName() string {
	return "t_user"
}
