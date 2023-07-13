package models

import "gorm.io/gorm"

type Permission struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	SystemRole   SystemRole
	SystemRoleID string
	Deleted      gorm.DeletedAt
}

func (u *Position) String() string {
	return u.Name
}
