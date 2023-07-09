package models

import "gorm.io/gorm"

type User struct {
	ID          uint `gorm:"primaryKey"`
	Username    string
	Password    string
	Name        string
	AccountType string
	Department  string
	Position    string
	Date        string
	Number      string
	Default     string
	Picture     string
	Deleted     gorm.DeletedAt
}

func (u *User) String() string {
	return u.Name
}
