package models

import "gorm.io/gorm"

type Position struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Description	 string
	Deleted      gorm.DeletedAt
}

func (u *Position) String() string {
	return u.Name
}


type Permissionlist struct {
	ID           uint `gorm:"primaryKey"`
	Position	 Position
	PositionID	 string
	Permission   Permission
	PermissionID string
	Deleted      gorm.DeletedAt
}
