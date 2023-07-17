package models

import "gorm.io/gorm"

type Location struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	Description    string
	Image          string
	ParentLocation string
	Code           string
	Category       string // remarks : Store Room, Plants, Building
	Deleted        gorm.DeletedAt
}
