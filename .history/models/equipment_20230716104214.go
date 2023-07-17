package models

import "gorm.io/gorm"

type Equipment struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Deleted gorm.DeletedAt
}
