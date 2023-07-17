package models

import "gorm.io/gorm"

type Location struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Deleted     gorm.DeletedAt