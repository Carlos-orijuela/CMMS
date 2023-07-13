package models

import "gorm.io/gorm"

type User struct {
	ID           uint `gorm:"primaryKey"`
	Username     string
	Email        string
	Password     string
	Name         string
	SystemRole   SystemRole
	SystemRoleID string
	Group        Group
	GroupID      string
	Date         string
	Number       string
	Deleted      gorm.DeletedAt
}

// func (u *User) String() string {
// 	return u.Name
// }
type SystemRole struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Permission  string
	Deleted     gorm.DeletedAt
}

type Group struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Deleted     gorm.DeletedAt
}
