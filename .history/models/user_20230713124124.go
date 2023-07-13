package models

import "gorm.io/gorm"

type User struct {
	ID           uint `gorm:"primaryKey"`
	Username     string
	Password     string
	Name         string
	AccountType  string
	Department   Department
	DepartmentID string
	Position     Position
	PositionID   string
	Date         string
	Number       string
	Deleted      gorm.DeletedAt
}

func (u *User) String() string {
	return u.Name
}

type SystemRole struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Deleted gorm.DeletedAt
}
