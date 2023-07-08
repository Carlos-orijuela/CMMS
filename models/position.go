package models

import "gorm.io/gorm"

type Position struct {
	ID          	uint `gorm:"primaryKey"`
	Name        	string
	Department		Department
	DepartmentID 	string
	Deleted     	gorm.DeletedAt
}

func (u *Position) String() string {
	return u.Name
}
