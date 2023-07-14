package models

import "gorm.io/gorm"

type User struct {
	ID           uint `gorm:"primaryKey"`
	Username     string
	Email        string
	Password     string
	Name         string
	Position  	 Position
	PositionID	 string
	Date         string
	Number       string
	Deleted      gorm.DeletedAt
}

//SYSTEM ROLE PERMISSION
type Permission struct {
	ID           uint `gorm:"primaryKey"`
	Description	 string
	Deleted      gorm.DeletedAt
}

type Group struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Deleted     gorm.DeletedAt
}
