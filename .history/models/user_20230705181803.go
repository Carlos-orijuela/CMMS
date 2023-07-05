package models

import "gorm.io/gorm"

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Password string
	Name     string
	Type     string
	LoginTime string
	OTP 	  string
	Date 	 string
	SerialNumber string
	Number	 string
	Code	 string
	Default  string
	Picture  string
	Deleted  gorm.DeletedAt
}

func (u *User) String() string {
	return u.Name
}
