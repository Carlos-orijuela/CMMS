package models

import "gorm.io/gorm"

type Department struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Deleted     gorm.DeletedAt
}

func (u *Department) String() string {
	return u.Name
}


type Grouplist struct {
	ID           uint `gorm:"primaryKey"`
	User	 	 User
	UserID	 	string
	Group   	Group
	GroupID		 string
	Deleted      gorm.DeletedAt
}

