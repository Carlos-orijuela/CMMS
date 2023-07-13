package models

import "gorm.io/gorm"

type Group struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Deleted gorm.DeletedAt
}

func (u *Group) String() string {
	return u.Name
}
