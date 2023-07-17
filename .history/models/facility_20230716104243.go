package models

import "gorm.io/gorm"

type Equipments struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Deleted gorm.DeletedAt
}
